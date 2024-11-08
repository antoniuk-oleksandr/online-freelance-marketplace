import {mantineTheme} from "@/utils/mantine-theme";
import {LayoutProps} from "@/types/LayoutProps";
import {MantineProvider} from "@mantine/core";
import {useTheme} from "next-themes";
import {useEffect, useState} from "react";

const MantineThemeProvider = (props: LayoutProps) => {
    const {children} = props;
    const [isMounted, setIsMounted] = useState(false);
    const {resolvedTheme} = useTheme();
    useEffect(() => setIsMounted(true), []);

    if (!isMounted) return null;
    return (
        <MantineProvider
            theme={{...mantineTheme(resolvedTheme === "dark" ? "dark" : "light")}}
        >
            {children}
        </MantineProvider>
    )
}

export default MantineThemeProvider;