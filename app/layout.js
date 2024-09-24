export const metadata = {
	title       : 'DummyAI',
	description : 'Artificial intelligence for dummies',
	manifest    : '/manifest.json',
	icons       : {
		apple : '/images/favicons/apple_touch_icon.png',
		icon  : [
			{
				url   : '/images/favicons/favicon.ico',
				sizes : 'any'
			},
			{
				url  : '/images/favicons/favicon.svg',
				type : 'image/svg+xml'
			}
		]
	}
};

export default function Layout({ children }) {

	return (
		<html lang = { 'en' }>
			<body>{ children }</body>
		</html>
	);
};
