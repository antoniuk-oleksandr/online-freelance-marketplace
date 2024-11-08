import {LayoutProps} from "@/types/LayoutProps";

const SubFormTextLayout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <div
            className={"flex justify-center gap-x-1 text-lg text-light-palette-text-secondary dark:text-dark-palette-text-secondary"}
        >
            {children}
        </div>
    )
}

export default SubFormTextLayout;