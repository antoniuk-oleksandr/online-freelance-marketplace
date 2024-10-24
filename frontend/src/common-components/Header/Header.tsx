import HeaderLayout from "./HeaderLayout";
import ThemeSwitch from "@/common-components/ThemeSwitch/ThemeSwitch";
import AppLogo from "@/common-components/AppLogo";

const Header = () => {
    return (
        <HeaderLayout>
            <AppLogo/>
            <ThemeSwitch/>
        </HeaderLayout>
    )
}

export default Header;