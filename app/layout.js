export const metadata = {
	title       : 'DummyAI',
	description : 'Artificial intelligence for dummies',
	manifest    : '/manifest.json',
	icons       : {
		icon  : [
			{
				url   : '/images/favicons/favicon.ico',
				sizes : 'any'
			},
			{
				url  : '/images/favicons/favicon.svg',
				type : 'image/svg+xml'
			}
		],
		apple : '/images/favicons/apple_touch_icon.png'
	}
};

export default function Layout({ children }) {

	return (
		<html lang = { 'en' }>
			<body>{ children }</body>
		</html>
	);
};
