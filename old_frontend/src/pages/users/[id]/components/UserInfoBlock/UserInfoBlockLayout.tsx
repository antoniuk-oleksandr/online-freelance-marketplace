import {LayoutProps} from "@/types/LayoutProps";

const Layout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <div className={"flex gap-4 items-center"}>
            {children}
        </div>
    )
}

export default Layout;