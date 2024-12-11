package main

import (
	"aoc"
	"fmt"
	"slices"
)

type block struct {
	id  int
	len int
}

func (b block) String() string {
	// var sb strings.Builder
	// for range b.len {
	// 	if b.id == -1 {
	// 		sb.WriteRune('.')
	// 	} else {
	// 		sb.WriteString(strconv.Itoa(b.id))
	// 	}
	// }
	// return sb.String()
	return fmt.Sprintf("{#%d, len: %d}", b.id, b.len)
}

func solve(input string) (int, int) {
	blocks := getFileBlocks(input)

	rIndex := len(blocks) - 1
	for rIndex%2 != 0 || blocks[rIndex].id == -1 {
		rIndex-- // If we have an empty block as the last one in the input it's not gonna move, ever
	}

	for rIndex > 0 {
		fileToMove := blocks[rIndex]
		if firstFree := findFirstFreeBlock(blocks, 0, rIndex); firstFree > 0 {
			blocks = defragFile(blocks, firstFree, fileToMove, &rIndex)
			// fmt.Println(blocks)
		}
		rIndex -= 2
	}
	res := getChecksum(blocks)

	blocks = getFileBlocks(input)
	rIndex = len(blocks) - 1
	// fmt.Println(blocks)
	for rIndex > 0 {
		fileToMove := blocks[rIndex]
		if fileToMove.id == -1 || fileToMove.len == 0 {
			rIndex--
			continue
		}
		if firstFree := findFirstFreeFittingBlock(blocks, rIndex, fileToMove.len); firstFree > 0 {
			//fmt.Println("Moving block ", fileToMove, "into", blocks[firstFree])
			blocks = defragFile(blocks, firstFree, fileToMove, &rIndex)

			//fmt.Println(blocks)
		}
		rIndex--
	}
	//fmt.Println(blocks)

	return res, getChecksum(blocks)
}

func findFirstFreeFittingBlock(blocks []*block, max int, size int) int {

	for i := 0; i < max; i++ {
		if blocks[i].id == -1 && blocks[i].len >= size {
			return i
		}
	}
	return -1
}

func getChecksum(blocks []*block) int {
	i := 0
	sum := 0
	for _, block := range blocks {
		if block.id == -1 {
			i += block.len
			continue
		}
		for m := i; m < i+block.len; m++ {
			sum += m * block.id
		}
		i += block.len
	}
	return sum
}

func defragFile(blocks []*block, targetIdx int, file *block, rIndex *int) []*block {
	targetBlock := blocks[targetIdx]
	if *rIndex < targetIdx {
		return blocks
	}
	if targetBlock.len == file.len {
		targetBlock.id = file.id
		file.id = -1
		return blocks
	}

	// If the target could fit more than the current file
	if targetBlock.len > file.len {
		remainingBlock := block{-1, targetBlock.len - file.len}
		targetBlock.len = file.len
		targetBlock.id = file.id
		file.id = -1
		*rIndex++ // increment right index as we're adding an extra element to the slice
		return slices.Insert(blocks, targetIdx+1, &remainingBlock)
	}

	// Target must be smaller than our file to handle
	targetBlock.id = file.id
	file.len -= targetBlock.len
	return defragFile(blocks, findFirstFreeBlock(blocks, targetIdx+1, len(blocks)), file, rIndex)
}

func getFileBlocks(input string) []*block {
	blocks := make([]*block, len(input))
	id := 0
	for i, r := range input {
		num := int(r - '0')
		var parsedFile block
		if i%2 == 0 { // Will remain in place
			parsedFile = block{id, num}
			id++
		} else {
			parsedFile = block{-1, num} // use ID -1 to denote empty spots on disk
		}
		blocks[i] = &parsedFile
	}

	return blocks
}

func findFirstFreeBlock(blocks []*block, min int, max int) int {
	for i := min; i < max; i++ {
		curr := blocks[i]
		if curr.id == -1 && curr.len > 0 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println("Example result:")
	fmt.Println(solve(example))

	fmt.Println("Real:")
	fmt.Println(solve(aoc.GetInputFromFile("09")))
}

const example = `2333133121414131402`
