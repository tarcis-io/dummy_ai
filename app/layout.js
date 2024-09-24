export const metadata = {
	title       : 'DummyAI',
	description : 'Artificial intelligence for dummies'
};

export default function Layout({ children }) {

	return (
		<html lang = { 'en' }>
			<body>{ children }</body>
		</html>
	);
};
