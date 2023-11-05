# I18N for Golang

提供关于时间的标准化多语言输出

## 目前支持的语言

| 语言| 缩写|FmtDuration示例 |FmtRelativeTime示例
| :-: | :-: | :-: |--|
| 简体中文 | zh_Hans| 1天6小时5分钟|1分钟前<p>明天
| 繁体中文 | zh_Hant| 1天6小時5分鐘|1分鐘前<p>明天
| 英语 | en| 1 day 6 hours 5 minutes|1 minute ago<p>tomorrow
| 日语 | ja| 1日6時間5分|1分前<p>明日
| 法语| fr| 1 jour 6 heures 5 minutes|il y a 1 minute<p>demain
| 西班牙语 | es| 1 día 6 horas 5 minutos|hace 1 minuto<p>mañana
| 葡萄牙语 | pt| 1 dia 6 horas 5 minutos|há 1 minuto<p>amanhã
| 德语| de| 1 Tag 6 Stunden 5 Minuten|vor 1 Minute<p>morgen
| 韩语/朝鲜语| ko| 1일6시간5분|1분 전<p>내일|
| 马来语| ms| 1 hari 6 jam 5 minit|1 minit lalu<p>hari ini|
| 泰语| th| 1 วัน 6 ชั่วโมง 5 นาที|1 นาทีที่ผ่านมา<p>พรุ่งนี้|
| 印尼语|id|1 hari 6 jam 5 menit|1 menit yang lalu<p>besok|
| 越南语|vi|1 ngày 6 giờ 5 phút|1 phút trước<p>Ngày mai|
| 俄语| ru| 1 день 6 часов 5 минут|1 минута назад<p>завтра|
| 阿拉伯语| ar| <div dir="rtl">1 يوم 6 ساعات 5 دقائق </div>|<div dir="rtl">1 دقيقة قبل</div> <div dir="rtl">غدًا</div> |

注意:阿拉伯语实在太复杂了，如果有需要强烈建议自己看源代码再校对一遍


## 更详细的示例参考网页
```
go run main/main.go
```
