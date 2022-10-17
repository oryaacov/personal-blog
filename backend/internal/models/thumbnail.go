package models

type Thumbnail struct {
	Id          string `json:"id" bson:"_id"`
	Title       string `json:"title" bson:"title"`
	Image       string `json:"img" bson:"image"`
	Summary     string `json:"summary" bson:"summary"`
	PublishedAt uint   `json:"date" bson:"publish_date"`
}

{"_id":"boot-file-using-grub-internal-hdd", "title":"How to boot from iso file using GRUB & internal HDD (without external USB/CD)", "image":"/public/thumbnails/grub.png", "summary":"I just received my new System76’s Lemp11 pro, and even though I love Pop!_OS I prefer to use Arch (btw), but… I didn’t have an external USB drive that will use me as installation media, so… Creating Partition  I have two internal hard drives and I’m going to shrink my secondary one to allocate more space for a new partition, you can skip this step in case you have some available memory.  Open Disks, and resize your partition  I Choose to leave 32GB my internal partition  in the free space create a new ext4 and name it as you want :)", "publish_date":1659733200000}