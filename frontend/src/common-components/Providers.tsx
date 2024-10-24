import {LayoutProps} from "@/types/LayoutProps";
import {ThemeProvider} from "next-themes";
import {MantineProvider} from "@mantine/core";

const Providers = (props: LayoutProps) => {
    const {children} = props;

    return (
        <ThemeProvider attribute="class" defaultTheme="system" enableSystem>
            <MantineProvider>
                {children}
            </MantineProvider>
        </ThemeProvider>
    )
}

export default Providers;