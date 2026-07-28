<!-- code-review-graph MCP tools -->
## MCP工具：code-review-graph

**重要提示：本项目搭载代码知识图谱。在使用 Grep/Glob/Read 遍历代码库之前，请务必优先调用 code-review-graph MCP 工具。**
图谱查询速度更快、Token消耗更低，并且能够提供文件扫描无法获取的结构化上下文（调用方、依赖项、测试覆盖情况等）。

### 哪些场景必须优先使用图谱工具
- **代码探索**：使用 `semantic_search_nodes_tool` 或 `query_graph_tool`，替代文本检索Grep
- **影响范围评估**：使用 `get_impact_radius_tool`，替代手动追踪导入依赖
- **代码评审**：`detect_changes_tool` + `get_review_context_tool`，避免读取全部文件
- **查询代码关联关系**：通过 `query_graph_tool` 查询调用者、被调用者、导入关系、关联测试用例
- **架构相关问题**：`get_architecture_overview_tool` + `list_communities_tool`

仅当图谱无法满足需求时，再降级使用 Grep/Glob/Read 文件读取工具。

### 核心工具清单

| 工具名称 | 使用场景 |
| ------ | ---------- |
| `detect_changes_tool` | 代码变更评审，输出带有风险评级的分析结果 |
| `get_review_context_tool` | 获取评审所需代码片段，节省Token |
| `get_impact_radius_tool` | 评估一处改动的影响辐射范围 |
| `get_affected_flows_tool` | 找出受改动影响的执行链路 |
| `query_graph_tool` | 追踪调用关系、导入依赖、关联测试等 |
| `semantic_search_nodes_tool` | 根据名称、关键词检索函数/类 |
| `get_architecture_overview_tool` | 获取代码库高层整体结构 |
| `refactor_tool` | 规划重命名、查找无效冗余代码 |

### 标准工作流程
1. 文件变更时图谱会通过钩子自动增量更新。
2. 执行代码评审先调用 `detect_changes_tool`。
3. 使用 `get_affected_flows_tool` 理清改动影响范围。
4. 使用 `query_graph_tool` 搭配 `tests_for` 查询关联测试覆盖情况。

## graphify

本项目知识图谱存放于 `graphify-out/`，包含聚合节点、模块群落结构以及跨文件依赖关系。

当用户输入 `/graphify --code-only`，执行任何操作前优先启用 graphify 相关能力。

规则：
- 询问代码库相关问题时，若存在 `graphify-out/graph.json`，优先执行 `graphify query "<question>" --code-only`；
  查询依赖链路使用 `graphify path "<A>" "<B>" --code-only`；定向概念解析使用 `graphify explain "<concept>" --code-only`。
  以上指令会返回范围收敛的子图谱，相比完整 GRAPH_REPORT.md 或全局文本检索体量更小。
- 钩子触发增量更新后，`graphify-out/` 内出现临时脏文件属于正常现象，不能以此为由跳过 graphify；
  只有任务目标是排查图谱过时、图谱数据错误，或是用户明确禁止使用图谱时，才可跳过。
- 如果存在 `graphify-out/wiki/index.md`，优先依靠该文档进行整体导航，而非直接浏览源码。
- 仅在全局架构评审、query/path/explain 无法获取足够信息时，再读取 `graphify-out/GRAPH_REPORT.md`。
- 修改代码完成后，执行 `graphify update . --code-only` 同步更新图谱（仅解析AST，无额外接口开销）。
- Monorepo Vue 项目中，出现节点数量为0的JSON文件、Vue节点去重相关警告属于正常现象，**不要视作构建错误**，正常继续使用图谱。
