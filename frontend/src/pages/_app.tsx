import type { AppProps } from "next/app"
import { ChakraProvider, extendTheme } from "@chakra-ui/react"
import "@fontsource/open-sans/400.css"

const theme = extendTheme({
    initialColorMode: "dark",
    useSystemColorMode: true,
    fonts: {
        heading: "-apple-system, Open Sans, sans-serif",
    },
})

export default function App({ Component, pageProps }: AppProps) {
    return (
        <ChakraProvider theme={theme}>
            <Component {...pageProps} />
        </ChakraProvider>
    )
}
