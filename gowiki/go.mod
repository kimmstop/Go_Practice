module wiki

replace Handler => ./Handler

replace Edit => ./Edit

replace Save => ./Save

replace View => ./View

replace Template => ./Template

replace Page => ./Page

go 1.15

require (
	Edit v0.0.0-00010101000000-000000000000
	Handler v0.0.0-00010101000000-000000000000
	Save v0.0.0-00010101000000-000000000000
	View v0.0.0-00010101000000-000000000000
)
