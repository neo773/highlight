import { LogLevel } from '@graph/schemas'
import { Preset } from '@highlight-run/ui'
import {
	backgroundColors,
	textColors,
} from '@highlight-run/ui/src/css/sprinkles.css'
import moment from 'moment'

export const FORMAT = 'YYYY-MM-DDTHH:mm:ss.000000000Z'

export const now = moment()
export const fifteenMinutesAgo = now.clone().subtract(15, 'minutes').toDate()
export const thirtyDaysAgo = now.clone().subtract(30, 'days').toDate()
export const PRESETS: Preset[] = [
	{
		startDate: fifteenMinutesAgo,
		label: 'Last 15 minutes',
	},
	{
		startDate: now.clone().subtract(60, 'minutes').toDate(),
		label: 'Last 60 minutes',
	},
	{
		startDate: now.clone().subtract(4, 'hours').toDate(),
		label: 'Last 4 hours',
	},
	{
		startDate: now.clone().subtract(24, 'hours').toDate(),
		label: 'Last 24 hours',
	},
	{
		startDate: now.clone().subtract(7, 'days').toDate(),
		label: 'Last 7 days',
	},
	{
		startDate: thirtyDaysAgo,
		label: 'Last 30 days',
	},
]

export const BACKGROUND_COLOR_MAPPING: {
	[key in LogLevel]: keyof typeof backgroundColors
} = {
	WARN: 'contentCaution',
	DEBUG: 'contentStrong',
	INFO: 'contentInformative',
	ERROR: 'contentBad',
	FATAL: 'contentBad',
	TRACE: 'contentStrong',
}

export const TEXT_COLOR_MAPPING: {
	[key in LogLevel]: keyof typeof textColors
} = {
	WARN: 'caution',
	DEBUG: 'strong',
	INFO: 'informative',
	ERROR: 'bad',
	FATAL: 'bad',
	TRACE: 'strong',
}

export type TIME_MODE = 'fixed-range' | 'permalink'