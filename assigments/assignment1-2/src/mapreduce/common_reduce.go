package mapreduce

import (
	"encoding/json"
	"os"
	"sort"
)

// doReduce does the job of a reduce worker: it reads the intermediate
// key/value pairs (produced by the map phase) for this task, sorts the
// intermediate key/value pairs by key, calls the user-defined reduce function
// (reduceF) for each key, and writes the output to disk.
func doReduce(
	jobName string, // the name of the whole MapReduce job
	reduceTaskNumber int, // which reduce task this is
	nMap int, // the number of map tasks that were run ("M" in the paper)
	reduceF func(key string, values []string) string,
) {
	// You will need to write this function.
	// You can find the intermediate file for this reduce task from map task number
	// m using reduceName(jobName, m, reduceTaskNumber).
	// Remember that you've encoded the values in the intermediate files, so you
	// will need to decode them. If you chose to use JSON, you can read out
	// multiple decoded values by creating a decoder, and then repeatedly calling
	// .Decode() on it until Decode() returns an error.
	//
	// You should write the reduced output in as JSON encoded KeyValue
	// objects to a file named mergeName(jobName, reduceTaskNumber). We require
	// you to use JSON here because that is what the merger than combines the
	// output from all the reduce tasks expects. There is nothing "special" about
	// JSON -- it is just the marshalling format we chose to use. It will look
	// something like this:
	//
	// enc := json.NewEncoder(mergeFile)
	// for key in ... {
	// 	enc.Encode(KeyValue{key, reduceF(...)})
	// }
	// file.Close()
	//
	// Use checkError to handle errors.

	// Read intermediate files from each map task
	var keyValues []KeyValue

	for m := 0; m < nMap; m++ {
		fileName := reduceName(jobName, m, reduceTaskNumber)

		file, err := os.Open(fileName)
		checkError(err)

		decoder := json.NewDecoder(file)

		for {
			var kv KeyValue

			err := decoder.Decode(&kv)

			if err != nil {
				break
			}

			keyValues = append(keyValues, kv)
		}
		file.Close()
	}

	// Sort keyValues by key
	sort.Slice(keyValues, func(i, j int) bool {
		return keyValues[i].Key < keyValues[j].Key
	})

	outFile, err := os.Create(mergeName(jobName, reduceTaskNumber))
	checkError(err)

	encoder := json.NewEncoder(outFile)

	// Group by key and apply reduceF
	i := 0
	for i < len(keyValues) {
		key := keyValues[i].Key

		var values []string

		for i < len(keyValues) && keyValues[i].Key == key {
			values = append(values, keyValues[i].Value)
			i++
		}

		// Call user-defined reduce function
		reducedValue := reduceF(key, values)

		err := encoder.Encode(&KeyValue{key, reducedValue})
		checkError(err)
	}

	outFile.Close()
}
