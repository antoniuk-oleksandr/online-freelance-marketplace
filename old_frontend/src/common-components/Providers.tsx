import {LayoutProps} from "@/types/LayoutProps";
import {ThemeProvider} from "next-themes";
import MantineThemeProvider from "@/common-components/MantineThemeProvider";
import {GoogleOAuthProvider} from "@react-oauth/google";

const Providers = (props: LayoutProps) => {
    const {children} = props;
    const clientId = "434086143424-g0o2v3g444htp3v4v8c1pvu3osp357g8.apps.googleusercontent.com";

    return (
        <ThemeProvider attribute="class" defaultTheme="system" enableSystem>
            <MantineThemeProvider>
                <GoogleOAuthProvider  clientId={clientId}>
                    {children}
                </GoogleOAuthProvider>
            </MantineThemeProvider>
        </ThemeProvider>
    )
}

export default Providers;