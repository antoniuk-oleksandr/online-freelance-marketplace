import "@/styles/globals.css";
import type {AppProps} from "next/app";
import Head from 'next/head';
import Header from "@/common-components/Header/Header";
import PageLayout from "@/common-components/PageLayout";
import Providers from "@/common-components/Providers";
import "@mantine/core/styles.css"
import {Inter} from "@next/font/google";
import Footer from "@/common-components/Footer/Footer";
import {usePage} from "@/general-hooks/use-page";
import {shouldHeaderBeShown} from "@/utils/utils";

const inter = Inter({
    subsets: ['latin'],
    weight: ['400', '500', '600', '700', '800', '900'],
});

export default function App({Component, pageProps}: AppProps) {
    const {page} = usePage();
    const headerShown = shouldHeaderBeShown(page);

    return (
        <Providers>
            <Head>
                <title>OFM | Online Freelance Marketplace</title>
                <meta name="description" content="Freelance marketplace for hiring and offering services online."/>
            </Head>
            <div className={`${inter.className} flex flex-col gap-y-8`}>
                {headerShown && <Header/>}
                <PageLayout>
                    <div className={headerShown ? 'min-h-app' : 'min-h-svh'}>
                        <Component {...pageProps} />
                    </div>
                </PageLayout>
                {headerShown && <Footer/>}
            </div>
        </Providers>
    )
}
