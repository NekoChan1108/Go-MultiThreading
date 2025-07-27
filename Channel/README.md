1. [windRegex](file://F:\GO-Projects\Go-MultiThreading\Channel\WindDirection.go#L12-L12):
    - 用于匹配完整的METAR气象报告中的风向信息
    - 匹配模式: `\d* METAR.*EGLL \d*Z [A-Z ]*(\d{5}KT|VRB\d{2}KT).*=`
    - 匹配示例: `200804270450 METAR COR EGLL 270450Z 14006KT 100V200 CAVOK 13/09 Q1015 NOSIG=`
    - 提取风向信息: `14006KT` 或 `VRB03KT` VRB开头表示风向不定 另一个的前三个表示风向角度

2. [tafValidation](file://F:\GO-Projects\Go-MultiThreading\Channel\WindDirection.go#L13-L13):
    - 用于检测是否包含TAF(终端机场预报)的行
    - 匹配模式: `.*TAF.*`
    - 匹配示例: 如果文件中有包含"TAF"的行

3. [comment](file://F:\GO-Projects\Go-MultiThreading\Channel\WindDirection.go#L14-L14):
    - 用于匹配注释行
    - 匹配模式: `\w*#.*`
    - 匹配示例: 以#号开头的注释行

4. [metarClose](file://F:\GO-Projects\Go-MultiThreading\Channel\WindDirection.go#L15-L15):
    - 用于匹配METAR报告结束标记
    - 匹配模式: `.*=`
    - 匹配示例: 所有以等号结尾的METAR报告行

5. [variableWind](file://F:\GO-Projects\Go-MultiThreading\Channel\WindDirection.go#L16-L16):
    - 用于匹配可变风向的报告
    - 匹配模式: `.*VRB\d{2}KT`
    - 匹配示例: `VRB03KT` (表示风向不定，风速3节)

6. [validWind](file://F:\GO-Projects\Go-MultiThreading\Channel\WindDirection.go#L17-L17):
    - 用于匹配标准格式的风向数据
    - 匹配模式: `\d{5}KT`
    - 匹配示例: `14006KT` (表示140度方向，风速6节)

7. [windDirOnly](file://F:\GO-Projects\Go-MultiThreading\Channel\WindDirection.go#L18-L18):
    - 用于提取固定的风向角度(前3位数字)
    - 匹配模式: `(\d{3})\d{2}KT`
    - 匹配示例: 从`14006KT`中提取`140`作为风向角度

8. [windDist](file://F:\GO-Projects\Go-MultiThreading\Channel\WindDirection.go#L19-L19):
    - 不是正则表达式，而是一个整型数组
    - 用于统计8个方向扇区的风向分布
    - 将360度分成8个扇区(每个45度)来统计风向频率