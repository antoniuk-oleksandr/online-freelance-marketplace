import "@/styles/globals.css";
import type {AppProps} from "next/app";
import Head from 'next/head';
import Header from "@/common-components/Header/Header";
import PageLayout from "@/common-components/PageLayout";
import Providers from "@/common-components/Providers";
import "@mantine/core/styles.css"
import {Lato} from "@next/font/google";
import Footer from "@/common-components/Footer/Footer";

const inter = Lato({
    subsets: ['latin'],
    weight: ['400', '700'],
});

export default function App({Component, pageProps}: AppProps) {
    return (
        <Providers>
            <Head>
                <title>OFM | Online Freelance Marketplace</title>
                <meta name="description" content="Freelance marketplace for hiring and offering services online."/>
            </Head>
            <div className={`${inter.className} flex flex-col gap-y-8`}>
                <Header/>
                <PageLayout>
                    <div className={"min-h-app"}>
                        <Component {...pageProps} />
                    </div>
                </PageLayout>
                <Footer/>
            </div>
        </Providers>
    )
}
