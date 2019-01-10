/*
“I Christopher Taliaferro (ch119541) affirm that this program is entirely my own work and that 
I have neither developed my code together with any another person, nor copied any code from any 
other person, nor permitted my code to be copied or otherwise used by any other person, nor 
have I copied, modified, or otherwise used programs created by others. I acknowledge that any 
violation of the above terms will be treated as academic dishonesty
*/

package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "math"
)

// Locate the index of a word from the input array
func getIndex(item string, list []string) (index int) {
  count := 0
  index = -1

  for i:= 0; i < len(list); i++ {
    if list[i] == item {
      index = count
      break
    }
    count++
  }
  return index
}

func printHeader(job string, lowerCYL int, upperCYL int, initCYL int, jobs [] int)() {
  fmt.Printf("Seek algorithm: %s\n", job)
  fmt.Printf("\tLower cylinder: %5d\n", lowerCYL)
  fmt.Printf("\tUpper cylinder: %5d\n", upperCYL)
  fmt.Printf("\tInit cylinder:  %5d\n", initCYL)
  fmt.Printf("\tCylinder requests:\n")

  for i := 0; i < len(jobs); i++{
    fmt.Printf("\t\tCylinder %5d\n", jobs[i])
  }
}

// Return MAX as 1st parameter and MIN as 2nd parameter
func hiLo(a int, b int)(int,int) {
  if a > b {
    return a, b
  } else {
    return b, a
  }
}

func bubblesort(items []int) ([] int) {
  var n = len(items)
  var sorted = false

  for !sorted {
    swapped := false
    for i := 0; i < n-1; i++ {
      if items[i] > items[i+1] {
        items[i+1], items[i] = items[i], items[i+1]
        swapped = true
      }
    }
    if !swapped {
      sorted = true
    }
    n = n - 1
  }
  return items
}

// Finds the closest number to 'val' in the 'options' array
// Returns that value and an array with that value removed
func closestValue(val int, options [] int) (int, [] int) {
  smallestDistance := math.MaxInt32
  var closestToVal int
  var closestToValIndex int

  for i := 0; i < len(options); i++ {
    max, min := hiLo(val, options[i])
    if (max - min) < smallestDistance {
      smallestDistance = max - min
      closestToVal = options[i]
      closestToValIndex = i
    }
  }
  options = append(options[:closestToValIndex], options[closestToValIndex+1:]...)
  return closestToVal, options
}

func firstComeFirstServe(lowerCYL int, upperCYL int, initCYL int, jobs [] int)() {
  printHeader("FCFS",lowerCYL, upperCYL, initCYL, jobs)
  var traversal int

  max, min := hiLo(jobs[0], initCYL)
  traversal += max - min

  for i:= 1; i < len(jobs); i++ {
    fmt.Printf("Servicing %5d\n", jobs[i-1])
    max, min := hiLo(jobs[i], jobs[i-1])
    traversal += max - min
  }
  fmt.Printf("Servicing %5d\n", jobs[len(jobs) - 1])
  fmt.Printf("FCFS traversal count = %5d\n", traversal)
}

func cLook(lowerCYL int, upperCYL int, initCYL int, jobs [] int)() {
  printHeader("C-LOOK",lowerCYL, upperCYL, initCYL, jobs)
  var traversal int
  index := 0
  jobs = bubblesort(jobs)

  // Determines what position in the array to begin
  for i:= 0; i < len(jobs); i++ {
    if initCYL > jobs[i] {
      index = i
      i++
    } else {
      if i != 0 {index++}
      break
    }
  }

  for i:= index; i < len(jobs); i++ {
    fmt.Printf("Servicing %5d\n", jobs[i])
  }

  // If the initial cylinder is greater then the smallest cylinder request
  // We need to the end and go back
  if initCYL > jobs[0] {
    for i:= 0; i < index; i++ {
      fmt.Printf("Servicing %5d\n", jobs[i])
    }
    traversal += (jobs[len(jobs)-1] - initCYL) + (jobs[len(jobs)-1] - jobs[0]) + (jobs[index - 1] - jobs[0])
  } else {
    traversal += jobs[len(jobs)-1] - initCYL
  }

  fmt.Printf("C-LOOK traversal count = %5d\n", traversal)
}

func look(lowerCYL int, upperCYL int, initCYL int, jobs [] int)() {
  printHeader("LOOK",lowerCYL, upperCYL, initCYL, jobs)
  var traversal int
  index := 0
  jobs = bubblesort(jobs)

  // Determines what position in the array to begin
  for i:= 0; i < len(jobs); i++ {
    if initCYL > jobs[i] {
      index = i
      i++
    } else {
      if i != 0 {index++}
      break
    }
  }

  for i:= index; i < len(jobs); i++ {
    fmt.Printf("Servicing %5d\n", jobs[i])
  }

  if initCYL > jobs[0] {
    traversal += jobs[len(jobs)-1] - initCYL
  } else {
    traversal += jobs[len(jobs)-1] - initCYL
  }

  if initCYL > jobs[0] {
    for i:= index - 1; i != -1; i-- {
      fmt.Printf("Servicing %5d\n", jobs[i])
    }
    traversal += jobs[len(jobs)-1] - jobs[0]
  }

  fmt.Printf("LOOK traversal count = %5d\n", traversal)
}

func cScan(lowerCYL int, upperCYL int, initCYL int, jobs [] int)() {
  printHeader("C-SCAN",lowerCYL, upperCYL, initCYL, jobs)
  var traversal int
  index := 0
  jobs = bubblesort(jobs)

  // Determines what position in the array to begin
  for i:= 0; i < len(jobs); i++ {
    if initCYL > jobs[i] {
      index = i
      i++
    } else {
      if i != 0 {index++}
      break
    }
  }

  for i:= index; i < len(jobs); i++ {
    fmt.Printf("Servicing %5d\n", jobs[i])
  }

  // If the initial cylinder is greater then the smallest cylinder request
  // We need to the end and go back
  if initCYL > jobs[0] {
    for i:= 0; i < index; i++ {
      fmt.Printf("Servicing %5d\n", jobs[i])
    }
    traversal += (upperCYL - initCYL) + (upperCYL - lowerCYL) + (jobs[index - 1] - lowerCYL)
  } else {
    traversal += jobs[len(jobs)-1] - initCYL
  }

  fmt.Printf("C-SCAN traversal count = %5d\n", traversal)
}

func scan(lowerCYL int, upperCYL int, initCYL int, jobs [] int)() {
  printHeader("SCAN",lowerCYL, upperCYL, initCYL, jobs)
  var traversal int
  index := 0
  jobs = bubblesort(jobs)

  // Determines what position in the array to begin
  for i:= 0; i < len(jobs); i++ {
    if initCYL > jobs[i] {
      index = i
      i++
    } else {
      if i != 0 {index++}
      break
    }
  }

  for i:= index; i < len(jobs); i++ {
    fmt.Printf("Servicing %5d\n", jobs[i])
  }

  if initCYL > jobs[0] {
    traversal += upperCYL - initCYL
  } else {
    traversal += jobs[len(jobs)-1] - initCYL
  }

  if initCYL > jobs[0] {
    for i:= index - 1; i != -1; i-- {
      fmt.Printf("Servicing %5d\n", jobs[i])
    }
    traversal += upperCYL - jobs[0]
  }

  fmt.Printf("SCAN traversal count = %5d\n", traversal)
}

func shortestSeekTimeFirst(lowerCYL int, upperCYL int, initCYL int, jobs [] int)() {
  printHeader("SSTF",lowerCYL, upperCYL, initCYL, jobs)
  var traversal int
  var jobsLeft [] int

  closestJob, jobsLeft := closestValue(initCYL, jobs)
  max, min := hiLo(closestJob, initCYL)
  traversal += max - min
  fmt.Printf("Servicing %5d\n", closestJob)

  for i := 0; len(jobsLeft) != 0 ;i++ {
    temp := closestJob
    closestJob, jobsLeft = closestValue(temp, jobsLeft)
    max, min := hiLo(closestJob, temp)
    traversal += max - min
    fmt.Printf("Servicing %5d\n", closestJob)
  }
  fmt.Printf("SSTF traversal count = %5d\n", traversal)
}

func main () {
  var data [] string
  var jobs [] int

  // Set up input and output file
  file := os.Args[1]
  input, _ := os.Open(file)
  defer input.Close()

  // Scanning each word
  scanner := bufio.NewScanner(input)
  scanner.Split(bufio.ScanWords)

  // data[] array will contain entire input file
  for scanner.Scan() {
    data = append(data, scanner.Text())
  }

  // Parse input file and initialize attributes
  index := getIndex("use", data)
  algorithm := data[index+1]

  index = getIndex("lowerCYL", data)
  lowerCYL, _ := strconv.Atoi(data[index+1])

  index = getIndex("upperCYL", data)
  upperCYL, _ := strconv.Atoi(data[index+1])

  index = getIndex("initCYL", data)
  initCYL, _ := strconv.Atoi(data[index+1])

  // Error Checking
  if upperCYL < lowerCYL {
    fmt.Printf("ABORT(13):upper (%d) < lower (%d)\n", upperCYL, lowerCYL)
    return
  } else if initCYL > upperCYL {
    fmt.Printf("ABORT(11):initial (%d) > upper (%d)\n", initCYL, upperCYL)
    return
  } else if initCYL < lowerCYL {
    fmt.Printf("ABORT(12):initial (%d) < lower (%d)\n", initCYL, lowerCYL)
    return
  }

  for i:= 0; i < len(data); i++ {
    if data[i] == "cylreq" {
      num, _ := strconv.Atoi(data[i + 1])
      if num > upperCYL || num < lowerCYL {
        fmt.Printf("ERROR(15):Request out of bounds: req (%d) > upper (%d) or < lower (%d)\n", num, upperCYL, lowerCYL)
      } else {
        jobs = append(jobs, num)
      }
    }
  }

  if algorithm == "fcfs" {
    firstComeFirstServe(lowerCYL, upperCYL, initCYL, jobs)
  } else if algorithm == "c-look" {
    cLook(lowerCYL, upperCYL, initCYL, jobs)
  } else if algorithm == "look" {
    look(lowerCYL, upperCYL, initCYL, jobs)
  } else if algorithm == "scan" {
    scan(lowerCYL, upperCYL, initCYL, jobs)
  } else if algorithm == "sstf" {
    shortestSeekTimeFirst(lowerCYL, upperCYL, initCYL, jobs)
  } else if algorithm == "c-scan" {
    cScan(lowerCYL, upperCYL, initCYL, jobs)
  }
}
