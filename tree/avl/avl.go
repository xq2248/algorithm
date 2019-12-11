/*
 * Copyright xq2248
 * mail: xq2248@163.com
 * github: xq2248
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package avl

//使用golang实现平衡二叉树，以及其操作和查找
//原理介绍图文并茂 https://www.sohu.com/a/270452030_478315

//二叉查找树
//1、若它的左子树不为空，则左子树上所有的节点值都小于它的根节点值。
//2、若它的右子树不为空，则右子树上所有的节点值均大于它的根节点值。
//3、它的左右子树也分别可以充当为二叉查找树。
//问题：极端情况下，左右子树偏重，导致查找时间复杂度编程O(n)

//AVL:平衡二叉树,特点
//1、具有二叉查找树的特点
//2、每个节点的左子树和右子树的高度相差不超过1

//S0为根节点，LeftOf(S0)=L1和R1位
//右旋:
//temp := r.left
//r.left = temp.right
//temp.right = r

//左旋:
//temp := r.right
//r.right = temp.left
//temp.left = r

//1、左-左型：做右旋。
//2、右-右型：做左旋。
//3、左-右型：先做左旋，后做右旋。
//4、右-左型：先做右旋，再做左旋。

type AVLTree struct {
	Data int   //节点的值
	Left,Right *AVLTree //左右子树
	Height int //树的高度
}

//LeftRotate 左旋
func (t *AVLTree) LeftRotate(){

}
//RightRotate 右旋
func (t *AVLTree) RightRotate(){
}

//

func (t *AVLTree) InsertData(){

}

func (t *AVLTree) DeleteData(){
}

func (t *AVLTree) Search(data int) bool{
	if t.Data == data {
		return true
	}
	return false
}