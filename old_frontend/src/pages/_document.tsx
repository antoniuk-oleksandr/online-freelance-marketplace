import {Html, Head, Main, NextScript} from "next/document";

export default function Document() {
    return (
        <Html lang="en">
            <Head/>
            <body
                className={"text-sm bg-light-palette-background-default dark:bg-dark-palette-background-default text-white-palette-text-primary dark:text-dark-palette-text-primary"}>
            <Main/>
            <NextScript/>
            </body>
        </Html>
);
}
