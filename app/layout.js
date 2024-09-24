'use client';

import {
	Page,
	PageSection
} from '@patternfly/react-core';

export default function Layout({ children }) {

	return (
		<html lang = { 'en' }>
			<body>
				<Page>
					<PageSection>{ children }</PageSection>
				</Page>
			</body>
		</html>
	);
};
