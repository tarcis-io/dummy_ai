export const metadata = {
	title       : 'DummyAI',
	description : 'Artificial intelligence for dummies',
	manifest    : '/manifest.json',
	icons       : {
		icon : {
			url  : '/images/favicons/favicon.svg',
			type : 'image/svg+xml'
		}
	}
};

export default function Layout({ children }) {

	return (
		<html>
			<body>{ children }</body>
		</html>
	);
};
