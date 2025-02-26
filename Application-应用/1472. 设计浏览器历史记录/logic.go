package main

/**
 * Created with GoLand 2022.2.3.
 * @author: zhashut
 * Date: 2025/2/26
 * Time: 20:15
 * Description: No Description
 */

type BrowserHistory struct {
	CurPage    string
	PrePages   []string
	CanForward int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		CurPage:  homepage,
		PrePages: []string{homepage},
	}
}

func (this *BrowserHistory) Visit(url string) {
	curIndex := this.getCurPageIndex()
	maxIndex := len(this.PrePages) - 1
	if curIndex == maxIndex {
		this.PrePages = append(this.PrePages, url)
	} else {
		this.PrePages[curIndex+1] = url
		this.PrePages = this.PrePages[:(curIndex+1)+1]
	}
	this.CurPage = url
	this.CanForward = 0
}

func (this *BrowserHistory) Back(steps int) string {
	// 找到当前 page 的下标，然后返回 -steps 个 url
	curIndex := this.getCurPageIndex() - steps
	if curIndex <= 0 {
		// 返回最前面的 url
		this.CanForward = len(this.PrePages) - 1
		this.CurPage = this.PrePages[0]
	} else {
		this.CanForward += steps
		this.CurPage = this.PrePages[curIndex]
	}
	return this.CurPage
}

func (this *BrowserHistory) Forward(steps int) string {
	curIndex := this.getCurPageIndex()
	if this.CanForward == 0 {
		return this.CurPage
	}
	if steps > this.CanForward {
		this.CurPage = this.PrePages[curIndex+this.CanForward]
		this.CanForward = 0
	} else {
		this.CurPage = this.PrePages[curIndex+steps]
		this.CanForward -= steps
	}

	return this.CurPage
}

func (this *BrowserHistory) getCurPageIndex() int {
	var maxCurIndex int
	for i, page := range this.PrePages {
		if this.CurPage == page && i >= maxCurIndex {
			maxCurIndex = i
		}
	}
	return maxCurIndex
}
