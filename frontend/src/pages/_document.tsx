import { ColorModeScript } from "@chakra-ui/react"
import NextDocument, { Html, Head, Main, NextScript } from "next/document"

const config = {}

export default function Document() {
    return (
        <Html lang="en">
            <Head>
                <link
                    rel="stylesheet"
                    href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.3/font/bootstrap-icons.css"
                />
            </Head>
            <body>
                {/* 👇 Here's the script */}
                <ColorModeScript initialColorMode={"system"} />
                <Main />
                <NextScript />
            </body>
        </Html>
    )
}
