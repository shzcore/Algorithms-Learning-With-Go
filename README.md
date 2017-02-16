算法学习 Golang 版 - 探索算法始末
=====================================

- [结构](#three)
- [算法](#four)
- [问题](#two)
- [原理](#one)
- [思想](#zero)

## <i id="three"></i>结构
- 组织结构
	- 线性结构
		- 线性表(Linear List)
		- 栈(Stack)
		- 队列(Queue)
		- 数组(Array)
		- 串(String)
		- 广义表(General List)
		- 跳跃表
	- 非线性结构
		- 树(Tree)
			- 胜者树
			- 左偏树
			- 二项树
			- 线段树
			- 后缀树
			- 哈夫曼树
			- trie树(静态建树、动态建树)
			- 静态二叉检索树
			- RMQ
			- 有限状态自动机理论
		- 堆(Heap)
			二叉堆
			斜堆
			
		- 图(Graph)
			trie图
- 统计结构
	- 树状数组
	- 虚二叉树
	- 线段树
	- 矩形面积并
	- 圆形面积并
- 关系结构
	- Hash表
	- 并查集
	- 路径压缩思想
	
## <i id="four"></i>算法

- 排序
	- 插入排序
		- 直接插入排序(Insert Sort) O(n^2)
		- 折半插入排序(Binary Insert Sort)
		- 希尔排序(Shell Sort)

	- 交换排序
		- 冒泡排序(Bubble Sort) O(n^2)
		- 快速排序(Quick Sort)?? O(nlogn)

	- 选择排序
		- 直接选择排序(Select Sort) O(n^2)
		- 锦标赛排序(Tournament Sort) O(nlogn)
		- 堆排序(Heap Sort) O(nlogn)

	- 归并排序(Merge Sort) O(nlogn)

	- 基数排序(Radix Sort) O(d(n+radix))

- 查找

	- 二分查找(Binary Search)

	- 树型
		- 二叉搜索树(Binary Search Tree)
		- 平衡搜索树(AVL Tree)
		- 并查集(Union-Find Set)

	- 哈希(Hashing)


- 以“相似或相同子问题”为核心的算法
	- 递推
	- 递归(Recursion)
	- 贪心算法(Greedy)
	- 动态规划

- 模拟算法

- 搜索算法
	- 最优化剪枝
	- 可行性剪枝
	- 记忆化搜索
	- 枚举搜索(Enumeration)
	- 深度优先(Depth First Search)
	- 广度优先(Breadth First Search)
	- 启发式搜索(Heuristic Search)		

- 机器学习算法
	- 回归算法
		- 最小二乘法(Ordinary Least Square)
		- 逻辑回归(Logistic Regression)
		- 逐步式回归(Stepwise Regression)
		- 多元自适应回归样条(Multivariate Adaptive Regression Splines)
		- 本地散点平滑估计(Locally Estimated Scatterplot Smoothing)
	- 基于实例的算法
		- k-Nearest Neighbor(KNN)
		- 学习矢量量化（Learning Vector Quantization， LVQ）
		- 自组织映射算法（Self-Organizing Map ， SOM）
	- 正则化方法 
		- Ridge Regression
		- Least Absolute Shrinkage and Selection Operator（LASSO）
		- 弹性网络（Elastic Net）
	- 决策树学习
		- 分类及回归树（Classification And Regression Tree， CART）
		- ID3 (Iterative Dichotomiser 3)
		- C4.5
		- Chi-squared Automatic Interaction Detection(CHAID)
		- Decision Stump
		- 随机森林（Random Forest）
		- 多元自适应回归样条（MARS）
		- 梯度推进机（Gradient Boosting Machine， GBM）
	- 贝叶斯方法
		- 朴素贝叶斯算法
		- 平均单依赖估计（Averaged One-Dependence Estimators， AODE）
		- Bayesian Belief Network（BBN）
	- 基于核的算法
		- 支持向量机（Support Vector Machine， SVM）
		- 径向基函数（Radial Basis Function ，RBF)
		- 线性判别分析（Linear Discriminate Analysis ，LDA)
	- 聚类算法
		- k-Means算法
		- 期望最大化算法（Expectation Maximization， EM）
	- 关联规则学习
		- Apriori算法
		- Eclat算法
	- 人工神经网络
		- 感知器神经网络（Perceptron Neural Network）
		- 反向传递（Back Propagation） 
		- Hopfield网络
		- 学习矢量量化（Learning Vector Quantization， LVQ）
	- 深度学习
		- 受限波尔兹曼机（Restricted Boltzmann Machine， RBN）
		- Deep Belief Networks（DBN）
		- 卷积网络（Convolutional Network）
		- 堆栈式自动编码器（Stacked Auto-encoders）
	- 降低维度算法
		- 主成份分析（Principle Component Analysis， PCA）
		- 偏最小二乘回归（Partial Least Square Regression，PLS）
		- Sammon映射
		- 多维尺度（Multi-Dimensional Scaling, MDS）
		- 投影追踪（Projection Pursuit）
	- 集成算法
		- Boosting
		- Bootstrapped Aggregation（Bagging）
		- AdaBoost
		- 堆叠泛化（Stacked Generalization， Blending）

## <i id="two"></i>问题

- 图论
	- 图的深度优先遍历和广度优先遍历.
	- 最短路径算法(dijkstra,bellman-ford,floyd,heap+dijkstra)
	- 最小生成树算法(prim,kruskal)
	- 拓扑排序
	- 二分图的最大匹配 (匈牙利算法)
	- 最大流的增广路算法(KM算法)
	- 度限制最小生成树和第K最短路.
	- 最优比率生成树.
	- 次小生成树.
	- 无向图、有向图的最小环
	- 差分约束系统的建立和求解
	- 最小费用最大流
	- 双连通分量
	- 强连通分支及其缩点
	- 图的割边和割点
	- 最小割模型、网络流规约

- 动态规划
	- 四边形不等式理论
	- 函数的凸凹性
	- 状态设计
	- 规划方向
	- 背包问题
	- 旅行商问题
	- 最优二分检索树
	- 树型动态规划
	- 状态动态规划
	- 记录状态的动态规划
	- LCA(Least Common Ancestors),即最近公共祖先
	- RMQ(Range Minimum/Maximum Query),即区间最值查询
	- 最长子序列系列问题
	- 最长不下降子序列
	- 最长公共子序列
	- 最长公共不下降子序列

## <i id="one"></i>原理

- 组合数学
	- 加法原理和乘法原理.
	- 排列组合
	- 容斥原理
	- 抽屉原理
	- 置换群与Polya定理
	- 递推关系和母函数
	- MoBius反演
	- 偏序关系理论
- 数论	
	- 素数与整除问题
	- 进制位
	- 同余模运算
- 博奕论	
	- 极大极小过程
	- Nim问题
- 概率论
- 数理统计
- 运筹学
- 计算几何学
	- 几何公式
	- 叉积和点积的运用(如线段相交的判定,点到线段的距离等)
	- 凸包
	- 坐标离散化.
	- 扫描线算法(例如求矩形的面积和周长并,常和线段树或堆一起使用).
	- 多边形的内核(半平面交)
	- 几何工具的综合应用
	- 半平面求交
	- 可视图的建立
	- 点集最小圆覆盖.
	- 对踵点
	

## <i id="zero"></i>思想
数学研究最抽象的数与结构的量、量变及其关系，算法用它与逻辑来发现从具象的问题空间到解空间的逻辑映射。
算法的思想就来自于数学、逻辑及求解的工具、方法与规律中。

附：[动态规划常见具体问题清单](https://github.com/ningskyer/Algorithms-Learning-With-Go/dp-problems.md)
