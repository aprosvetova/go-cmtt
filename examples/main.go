package main

import (
	"fmt"
	"github.com/aprosvetova/go-cmtt"
	"log"
	"sort"
	"time"
)

func main() {
	client, err := cmtt.NewWithQr("<YOUR_TOKEN_FROM_QR>", "tjournal")
	if err != nil {
		log.Fatal(err)
	}
	commentIDs := getCommentIDs(client)
	names, likes, dislikes := getAllLikers(client, commentIDs)
	fmt.Println("Likes:")
	for _, p := range likes {
		fmt.Println(names[p.key], p.value)
	}
	fmt.Println("\nDislikes:")
	for _, p := range dislikes {
		fmt.Println(names[p.key], p.value)
	}
}

func getCommentIDs(client *cmtt.Cmtt) []int {
	var commentIDs []int
	offset := 0
	for {
		res, err := client.GetMineComments(50, offset)
		offset += 50
		if err != nil {
			log.Fatal(err)
		}
		for _, comment := range *res {
			commentIDs = append(commentIDs, comment.Id)
		}
		if len(*res) < 50 {
			break
		}
	}
	return commentIDs
}

func getAllLikers(client *cmtt.Cmtt, commentIDs []int) (map[int]string, pairList, pairList) {
	names := make(map[int]string)
	likes := make(map[int]int)
	dislikes := make(map[int]int)
	for _, comment := range commentIDs {
		likers, err := getLikers(client, comment)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Millisecond * 1000)
		for liker, info := range *likers {
			names[liker] = info.Name
			if info.Sign == 1 {
				likes[liker]++
			} else {
				dislikes[liker]++
			}
		}
	}
	return names, orderMapDescending(likes), orderMapDescending(dislikes)
}

func getLikers(client *cmtt.Cmtt, commentID int) (*map[int]cmtt.Liker, error) {
	likers, err := client.GetCommentLikers(commentID)
	if err != nil {
		if err.Error() == "timeout" {
			time.Sleep(5 * time.Second)
			return getLikers(client, commentID)
		} else {
			return nil, err
		}
	}
	return likers, nil
}

func orderMapDescending(m map[int]int) pairList {
	p := *new(pairList)
	for k, v := range m {
		p = append(p, pair{k, v})
	}
	sort.Sort(sort.Reverse(p))
	return p
}

type pair struct {
	key   int
	value int
}

type pairList []pair

func (p pairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p pairList) Len() int           { return len(p) }
func (p pairList) Less(i, j int) bool { return p[i].value < p[j].value }
