import {LayoutProps} from "@/types/LayoutProps";

const PackageTabListLayout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <div className={"flex"}>
            {children}
        </div>
    )
}

export default PackageTabListLayout;