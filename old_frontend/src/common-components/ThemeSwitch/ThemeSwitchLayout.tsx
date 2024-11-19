import {LayoutProps} from "@/types/LayoutProps";

const ThemeSwitchLayout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <div>
            {children}
        </div>
    )
}

export default ThemeSwitchLayout;