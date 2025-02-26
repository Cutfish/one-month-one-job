### 1、快慢指针找环
为什么快指针和慢指针总是会在环中某一点相遇，而不会刚好“跳过”彼此？
由于快指针每次跳两步，慢指针每次跳一步，快指针实际上比慢指针多走了一步（两步-一步）。每当快指针和慢指针没有相遇时，快指针相当于以慢指针的速度“追”慢指针。
```C++
if(!head || !head->next) return false;
ListNode *fast = head, *low = head;
while (fast && fast->next) {
    fast = fast->next;
    fast = fast->next;
    low = low->next;
    if(fast == low ) return true; 
}
return false;
```
### 2、两数之和
1. 排序、左右指针、向中间移动。根据大小关系左指右移或者右指左移。
2. hash.
### 3、两数相加（链表）
listNode val carry.
### 4、无重复字符的最长子串
贪心、hash、左右指针。右指针往右走，一旦遇到重复字符，左指针右移动hash值减一，判断最大值。
### 5、两个正序数组的中位数
```C++
if (left < m && (right >= n || nums1[left] <= nums2[right])) {
    curr = nums1[left++];
} else {
    curr = nums2[right++];
}
``` 
### 6、最长回文子串
1. 暴力（n*3）不行
---
2. 动态规划
第一层循环用len，第二层循环用左端点

### 7、正则表达式匹配
```C++
// 初始化f[0][j]，处理模式以 '*' 开头的情况
for (int j = 2; j <= n; j++) {
    if (p[j] == '*' && f[0][j - 2]) {
        f[0][j] = true;  // '*' 匹配空字符串
    }
}

//处理f[i][j]
if (p[j] == s[i] || p[j] == '.') {
    f[i][j] = f[i - 1][j - 1];  // 匹配单个字符或 '.' 
} else if (p[j] == '*') {
    // '*' 匹配 0 或多个字符
    f[i][j] = f[i][j - 2] ||  // '*' 匹配 0 次
        f[i - 1][j - 2] && (p[j - 1] == s[i] || p[j - 1] == '.') ||
        f[i - 1][j] && (p[j - 1] == s[i] || p[j - 1] == '.');
        // 前面两行可以缩写为：(f[i - 1][j] && (p[j - 1] == s[i] || p[j - 1] == '.'));  // '*' 匹配多个
}
```

### 8、盛最多水的容器
双指针向中间靠拢，谁低谁移动

### 9、三数之和
1. 排序
2. 一次循环+双指针（在循环右侧） 
3. 符合条件时。若下一个是重复，一直移动到最后一个重复的。两个指针都动一下

### 10、电话号码的数字组合
```C++
void dfs(const string& digits, int cur ,string word, vector<string>& ret,const vector<string>& list) {
    if (cur == digits.size()) ret.push_back(word);
    else {
        int index = digits[cur] - '0';
        for (int i = 0; i < list[index].size(); i++) {
            dfs(digits, cur + 1,word + list[index][i], ret, list);
        }
    }
}

dfs(digits, 0 ,"", ret, list);


```
### 11、删除倒数第n个节点
1. 快慢指针间隔n
2. 计算长度减去n
### 12、合并两个有序链表
建新节点取小节点
### 13、括号生成
1. 终止：L == 0 && R == 0
2. 任何时候：L <= R
---
```C++
//风格一
string buf;
void dfs(int left, int right) {
    if (left == 0 && right == 0) {
        ret.push_back(buf);return;
    }
    if (left > 0) {
        buf.push_back('(');dfs(left - 1, right);buf.pop_back();
    }
    if (left < right) {
        buf.push_back(')');dfs(left, right - 1);buf.pop_back();
    }
}

//风格二
char buf[20];
void dfs(int left, int right, int cur) {
    if ( !left && !right ) {
        ret.push_back(buf);return;
    }
    if (left ) {
        buf[cur] = '(';dfs(left - 1, right, cur + 1);
    }
    if (left < right) {
        buf[cur] = ')';dfs(left, right - 1,cur + 1 );
    }
}

```

### 14、合并K个升序链表
- 优先队列取链表数组的每一个链表的头
- while (!pq.empty())
- 拿走top之后，放该节点的下一个节点
```C++
typedef pair<int, ListNode*> Elem;
using Elem = pair<int, ListNode*>;

priority_queue<Elem, vector<Elem>, greater<Elem>> pq;
//greater是小根堆，取出来是最小值 less...
```

### 15、下一个排列
1. 从右到左，找到第一个a[i] < a[j]，若没有就直接反转整个vector
2. 从右往左找到第一个a[j] > a[i]（至少有一个a[j+1]）
3. 交换a[i]和a[j]，然后反转i + 1到最后一个元素
### 16、最长有效括号
1. 栈里面存括号的index，提前存一个-1
2. 遇到“（”就入栈。遇到“）”判断栈顶是不是-1或者栈顶是“）”也入栈，否则就pop，然后max(ret, i-stk.top())
### 17、在排序数组中查找元素的第一个和最后一个位置
```C++
int bin1(const vector<int>& nums, int &target) {
    int left = 0, right = nums.size() - 1; 
    int ans = -1;
    while (left <= right) {
        int mid = (left + right ) >> 1;
        if (nums[mid] < target) left = mid + 1;
        else  {
            ans = mid;
            right = mid - 1;    
        }
    }
    if (ans == -1 || nums[ans] != target) return -1;
    return ans;
}
int bin2(const vector<int>& nums, int &target) {
    int left = 0, right = nums.size() - 1; 
    int ans = -1;
    while (left <= right) {
        int mid = (left + right ) >> 1;
        if (nums[mid] > target) right = mid - 1; 
        else  {
            ans = mid;
            left = mid + 1;
        }
    }
    if (ans == -1 || nums[ans] != target) return -1;
    return ans;
}
```
### 17、只出现一次的数字
位运算 ^ (异或) 相同为0，不同为1
```
0^a^a^b^b^...^x
=0^0^...^x
=x
```
满足交换律，顺序无所谓

### 18、搜索插入的位置


### 19、回文链表
1. 复制出来新的链表（倒置）
2. 快慢指针，慢指针到中间节点，转置后半链表

### 20、倒置链表
```C++
while(head) {
    newHead = head;
    newHead->next = p;
    p = newHead;
    head = head->next;  //注意此项的位置，不能在第二行的上面，否则就是超时（循环）
}
```
### 21、轮转数组
- 和交换前几个元素和后几个元素的位置的题的区别在于要取模
- reverse(a, b) 左闭右开
### 22、合并区间
1. 按照左端点排序
2. 更新L和R
### 23、寻找重复数
nums中的数作为index，使index位置的数字发生变化（变负数或者大于n的数字）
### 24、颜色分类
1. 循环两遍
2. 双指针
```C++
//注意看对于指针指向2的情况，还有终止条件
while (i <= right) {
    if (nums[i] == 0) {
        swap(nums[i], nums[left]);
        i++;left++;
    }else if (nums[i] == 2){
        swap(nums[i], nums[right]);right--;
    }
    else i++;
}
```
### 25、最小栈
1. 从零开始，可以注意实现扩展数组
2. 用两个栈实现，一个正常，另外一个有可能装当前数字（栈顶数字一定是最小的），如果等于栈顶数字也装进去
### 26、最大子序和
- 如果sum > 0，则说明 sum 对结果有增益效果，则 sum 保留并加上当前遍历数字
- 如果sum <= 0，则说明 sum 对结果无增益效果，需要舍弃，则 sum 直接更新为当前遍历数字
- 每次比较sum和ans的大小，将最大值置为ans，遍历结束返回结果

### 27、树的直径
其实就是调用求树的深度代码，在里面加一个ret = max(ret, left + right)
### 28、对称二叉树
很多时候一个递归函数是不用的，写一个额外的递归函数，再在主函数里面调用就可以了
```C++
bool check(TreeNode* left, TreeNode* right) {
    if (!left && !right) return true;
    if (!left || !right) return false;
    return left->val == right->val && check(left->left, right->right) && check(left->right, right->left);
}
bool isSymmetric(TreeNode* root) {
    return check(root->left, root->right);
}
```
### 29、将有序数组转换为二叉搜索树
- 中序搜索
- 取数组index的中间值作为根节点
- 一个主函数调用一个递归函数
### 30、验证二叉搜索树
- 利用的是中序遍历的所有值是递增的
```C++
long long INF_MIN = LONG_MIN;
bool isValidBST(TreeNode* root) {
    if (!root) return true;
    bool l = isValidBST(root->left);
    if (root->val <= INF_MIN) return false;
    INF_MIN = root->val;
    return l && isValidBST(root->right);
}
```
### 31、BST中第k小的元素
和30类似，用中序遍历，更新count
### 32、组合总和
```C++
class Solution {
public:
    vector<vector<int>> ret;
    vector<int> b;

    void dfs(vector<int> &candidates, int p, int sum, int target) {
        if (p == candidates.size()) {
            return;
        }
        if (sum == target) {
            ret.push_back(b);
            return;
        }
        dfs(candidates, p + 1, sum, target);

        if (sum + candidates[p] <= target) {
            b.push_back(candidates[p]);
            dfs(candidates, p, sum + candidates[p], target);
            b.pop_back();
        }

    }
    vector<vector<int>> combinationSum(vector<int>& candidates, int target) {
        dfs(candidates, 0, 0, target);
        return ret;
    }
};
```
### 33、字母异位词分组
- 注意重复的也要
- map<string vector<string>>
- 排完序当key,原值放value中
### 34、跳跃游戏
1. 左指针记录当前位置，右指针是最远能到的位置，两指针相遇代表不行
2. dp: bool f[i] 表示从这个位置能否到末尾，从右往左遍历
### 35、不同路径
dp:f[i][j]表示f[0]f[0]到当前位置的路径数量 **(f[1][1] = 1)**
### 36、最小路径和
与35类似,区别在于如何更新（之前都是加1，现在是加位置上面的数字）
### 37、编辑距离
```C++
if (word1[i-1] == word2[j-1]) f[i][j] = f[i-1][j-1]; 
else f[i][j] = f[i-1][j-1] + 1;
f[i][j] = min(f[i][j],f[i-1][j] +1);
f[i][j] = min(f[i][j], f[i][j-1] +1);
```
### 38、二叉树展开为链表(不太能理解为什么要备份)
```C++
TreeNode* pre;
void dfs(TreeNode* root) {
    if(!root) return;

    pre->right = root;
    pre->left = nullptr;

    TreeNode* left = root->left;
    TreeNode* right = root->right;
    
    pre = root;
    dfs(left);
    dfs(right);
}
```
### 39、二叉树的层序遍历
qu不为空就遍历，获取当前size，for遍历输入当前值到vector、输入孩子到qu
### 40、岛屿数量
深度搜索，走过的改值，扫描上下左右
### 41、实现前缀树（Trie）
26叉树
```C++
class Trie {
private:
    bool isEnd;
    Trie* next[26];
public:
    Trie() {
        this->isEnd = false;
        for (int i= 0;i < 26; i++) {
            this->next[i] = nullptr;
        }
    }

    void insert(string word) {
        Trie* node = this;
        for (char c : word) {
            if (node->next[c - 'a'] == nullptr) node->next[c - 'a'] = new Trie();
            node = node->next[c - 'a'];
        }
        node->isEnd = true;
    }
    
    bool search(string word) {
        Trie* node = this;
        for (char c : word) {
            node = node->next[c - 'a'];
            if (node == nullptr) {
                return false;
            }
        }
        return node->isEnd;
    }
```
### 42、输出链表中倒数k的节点及其链表
```C++
if (!pHead) return nullptr;
ListNode *fast = pHead, *slow = pHead;
while (k--) {
    if (fast == nullptr) return nullptr;
    fast = fast->next;
}

while (fast) {
    fast = fast->next;
    slow = slow->next;
}
return slow;
```
### 43、合并二叉树
```C++
TreeNode* mergeTrees(TreeNode* t1, TreeNode* t2) {
    // write code here
    if (!t1) return t2;
    if (!t2 ) return t1;
    t1->val += t2->val;
    t1->left = mergeTrees(t1->left, t2->left);
    t1->right = mergeTrees(t1->right, t2->right);
    return t1;
}
```
### 44、字符串变形
- 将字符串大小写反转
- 将整个字符串的所有单词位置反转
**思路：**
1. stack
2. 先都反转，然后按照遇到空格反转单词 （空间复杂度常数级）
### 45、最长公共前缀
- 注意外层的for循环是第一个字符串的长度

### 46、最长连续序列
- 用set<int>
- string.substr(a,len)，第二个参数是字串长度（此题可不用这个函数）
### 47、找到字符串中所有字母异位词
- 滑动窗口
- 构造26长度的vector<int> 这个容器重载了==运算符，可以直接用
### 48、缺失的第一个正数
1. 数组长度为N，那么缺失的第一个数字一定在[1, N+1]中。
并且所有的数字是1 ~ N的时候，第一个缺失的正数才是N + 1
2. 遍历数组，所有小于等于0的数字变为 N + 1
3. 遍历数组，对于数字a，若abs(a)小于等于N，将nums[a-1] = -abs(nums[a-1])
4. 遍历数组，对于第一个出现的正数，返回其下标+1，若没有就返回N+1
### 49、除自身以外数组的乘积
**方法**
1. 两个辅助数组一个结果数组，辅助数组各自维护该位置数组元素的左边乘积和右边乘积
2. 只有一个结果数组，第一遍循环保存该位置对应数字的左边乘积，第二遍循环借助辅助变量R = 1，辅助数组元素 * R，然后用R保存当前位置右边乘积
### 50、和为 K 的子数组
1. 暴力循环
2. 哈希表
```C++
class Solution {
public:
    int subarraySum(vector<int>& nums, int k) {
        unordered_map<int, int> mp;
        mp[0] = 1;
        int count = 0, pre = 0;
        for (auto& x:nums) {
            pre += x;
            if (mp.find(pre - k) != mp.end()) {
                count += mp[pre - k];
            }
            mp[pre]++;
        }
        return count;
    }
};
```
### 51、矩阵置0
1. 借助标记数组或者set存标志行列：T(n) = O(mn),S(n) = O(m + n)
2. 直接将标志位放到第一行第一列，S(n) = O(1)
### 52、
### 53、
### 54、
### 55、
### 56、





​
 










