{{ define "index" }}
<!DOCTYPE html>
<html>
	<head>
		<title>Home</title>
	</head>
	<body>
        <p>Available locales:</p>
        <ul>
        {{range .Locales}} 
            <li><a href="/?locale={{.}}">{{.}}</a></li>
        {{end}}
        </ul>
		<p>Locale: {{ .Trans.Locale }}</p>
        <p>FmtNumber Positive: {{ .Trans.FmtNumber .PositiveNum 2 }}</p>
        <p>FmtNumber Negative: {{ .Trans.FmtNumber .NegativeNum 2 }}</p>
        <p>FmtPercent Negative: {{ .Trans.FmtPercent .Percent 2 }}</p>
        <p>FmtCurrency Negative: {{ .Trans.FmtCurrency .PositiveNum 2 .Trans.Currency }}</p>
        <p>FmtCurrency Negative: {{ .Trans.FmtCurrency .NegativeNum 2 .Trans.Currency }}</p>
        <p>FmtAccounting Negative: {{ .Trans.FmtAccounting .PositiveNum 2 .Trans.Currency }}</p>
        <p>FmtAccounting Negative: {{ .Trans.FmtAccounting .NegativeNum 2 .Trans.Currency }}</p>
        <p>FmtDateShort: {{ .Trans.FmtDateShort .Now }}</p>
        <p>FmtDateMedium: {{ .Trans.FmtDateMedium .Now }}</p>
        <p>FmtDateLong: {{ .Trans.FmtDateLong .Now }}</p>
        <p>FmtDateFull: {{ .Trans.FmtDateFull .Now }}</p>
        <p>FmtTimeShort: {{ .Trans.FmtTimeShort .Now }}</p>
        <p>FmtTimeMedium: {{ .Trans.FmtTimeMedium .Now }}</p>
        <p>FmtTimeLong: {{ .Trans.FmtTimeLong .Now }}</p>
        <p>FmtTimeFull: {{ .Trans.FmtTimeFull .Now }}</p>
        <p>FmtRelativeTime: {{ .Trans.FmtRelativeTime 0 0 }}</p>
        <p>FmtRelativeTime: {{ .Trans.FmtRelativeTime 10 0 }}</p>
        <p>FmtRelativeTime: {{ .Trans.FmtRelativeTime -1 1 }}</p>
        <p>FmtRelativeTime: {{ .Trans.FmtRelativeTime 0 1 }}</p>
        <p>FmtRelativeTime: {{ .Trans.FmtRelativeTime 0 2 }}</p>
        <p>FmtRelativeTime: {{ .Trans.FmtRelativeTime 1 2 }}</p>
        <p>FmtRelativeTime: {{ .Trans.FmtRelativeTime -1 3 }}</p>
        <p>FmtRelativeTime: {{ .Trans.FmtRelativeTime 0 3 }}</p>
        <p>FmtRelativeTime: {{ .Trans.FmtRelativeTime 1 3 }}</p>
        <p>FmtRelativeTime: {{ .Trans.FmtRelativeTime -1 4 }}</p>
        <p>FmtRelativeTime: {{ .Trans.FmtRelativeTime 0 4 }}</p>
        <p>FmtRelativeTime: {{ .Trans.FmtRelativeTime 1 4 }}</p>
        <p>FmtRelativeTime: {{ .Trans.FmtRelativeTime 2 4 }}</p>
        <p>FmtDuration: {{ .Trans.FmtDuration 50 }}</p>
        <p>FmtDuration: {{ .Trans.FmtDuration 90 }}</p>
        <p>FmtDuration: {{ .Trans.FmtDuration 3780 }}</p>
        <p>FmtDuration: {{ .Trans.FmtDuration 7505 }}</p>
        <p>FmtDuration: {{ .Trans.FmtDuration 108300 }}</p>
        <p>MonthsAbbreviated: {{ .Trans.MonthsAbbreviated }}</p>
        <p>MonthsNarrow: {{ .Trans.MonthsNarrow }}</p>
        <p>MonthsWide: {{ .Trans.MonthsWide }}</p>
        <p>WeekdaysAbbreviated: {{ .Trans.WeekdaysAbbreviated }}</p>
        <p>WeekdaysNarrow: {{ .Trans.WeekdaysNarrow }}</p>
        <p>WeekdaysShort: {{ .Trans.WeekdaysShort }}</p>
        <p>WeekdaysWide: {{ .Trans.WeekdaysWide }}</p>
	</body>
</html>
{{ end }}