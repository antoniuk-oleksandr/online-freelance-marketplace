import {LayoutProps} from "@/types/LayoutProps";

const SelectedPackageLayout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <div className={"p-6 text-base pt-8 text-light-palette-text-secondary dark:text-dark-palette-text-secondary"}>
            {children}
        </div>
    )
}

export default SelectedPackageLayout;