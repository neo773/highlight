import React from 'react'

import { IconProps } from './types'

export const IconSolidTraces = ({ size = '1em', ...props }: IconProps) => {
	return (
		<svg
			xmlns="http://www.w3.org/2000/svg"
			width={size}
			height={size}
			fill="none"
			viewBox="0 0 16 14"
			focusable="false"
			{...props}
		>
			<path
				fill="currentColor"
				fillRule="evenodd"
				clipRule="evenodd"
				d="M1.3999 3.49998C1.3999 2.72678 2.0267 2.09998 2.7999 2.09998H11.1999C11.9731 2.09998 12.5999 2.72678 12.5999 3.49998C12.5999 4.27317 11.9731 4.89998 11.1999 4.89998H2.7999C2.0267 4.89998 1.3999 4.27317 1.3999 3.49998Z"
			/>
			<path
				fill="currentColor"
				fillRule="evenodd"
				clipRule="evenodd"
				d="M1.3999 6.99998C1.3999 6.22678 2.0267 5.59998 2.7999 5.59998H4.1999C4.9731 5.59998 5.5999 6.22678 5.5999 6.99998C5.5999 7.77317 4.9731 8.39997 4.1999 8.39997H2.7999C2.0267 8.39997 1.3999 7.77317 1.3999 6.99998Z"
			/>
			<path
				fill="currentColor"
				fillRule="evenodd"
				clipRule="evenodd"
				d="M6.2999 6.99998C6.2999 6.22678 6.9267 5.59998 7.6999 5.59998C8.4731 5.59998 9.0999 6.22678 9.0999 6.99998C9.0999 7.77317 8.4731 8.39997 7.6999 8.39997C6.9267 8.39997 6.2999 7.77317 6.2999 6.99998Z"
			/>
			<path
				fill="currentColor"
				fillRule="evenodd"
				clipRule="evenodd"
				d="M9.7999 6.29998C9.7999 5.91338 10.1133 5.59998 10.4999 5.59998C10.8865 5.59998 11.1999 5.91338 11.1999 6.29998V7.69998C11.1999 8.08657 10.8865 8.39997 10.4999 8.39997C10.1133 8.39997 9.7999 8.08657 9.7999 7.69998V6.29998Z"
			/>
			<path
				fill="currentColor"
				fillRule="evenodd"
				clipRule="evenodd"
				d="M3.1499 10.5C3.1499 9.72678 3.7767 9.09997 4.5499 9.09997H9.4499C10.2231 9.09997 10.8499 9.72678 10.8499 10.5C10.8499 11.2732 10.2231 11.9 9.4499 11.9H4.5499C3.7767 11.9 3.1499 11.2732 3.1499 10.5Z"
			/>
		</svg>
	)
}
