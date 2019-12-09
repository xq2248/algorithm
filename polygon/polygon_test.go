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
package polygon

import (
	"testing"
)

func TestWalk(t *testing.T) {
	var sixPoints []Point = []Point{
		{1.0, 1.0},
		{1.0, 2.0},
		{3.0, 2.0},
		{3.0, 1.0},
		//{6.0, 4.0},
		//{6.0, 1.0},
	}
	Walk(sixPoints, 3)
}
