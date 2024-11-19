import ThemeSwitchLayout from "./ThemeSwitchLayout";
import {useEffect, useState} from "react";
import {useTheme} from "next-themes";
import {FaMoon, FaSun} from "react-icons/fa6";

const ThemeSwitch = () => {
    const [mounted, setMounted] = useState(false);
    const {setTheme, resolvedTheme} = useTheme();

    useEffect(() => setMounted(true), []);

    if (!mounted) return <div></div>;

    return (
        <ThemeSwitchLayout>
            <div
                className={"text-lg flex items-center cursor-pointer duration-200 ease-out"}
                onClick={() => setTheme(resolvedTheme === 'dark' ? 'light' : 'dark')}
            >
                {resolvedTheme === 'dark' ? <FaSun/> : <FaMoon/>}
            </div>
        </ThemeSwitchLayout>
    );
}

export default ThemeSwitch;