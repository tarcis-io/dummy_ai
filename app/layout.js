export const metadata = {
	title       : 'DummyAI',
	description : 'Artificial intelligence for dummies',
	manifest    : 'manifest.json'
};

export default function Layout({ children }) {

	return (
		<html>
			<body>{ children }</body>
		</html>
	);
};
