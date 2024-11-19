import {LayoutProps} from "@/types/LayoutProps";
import {Children} from "react";

const LeftSignSideOrTextLayout = (props: LayoutProps) => {
    const {children} = props;
    const childrenArray = Children.toArray(children);

    return (
        <div className={"relative text-light-palette-text-secondary dark:text-dark-palette-text-secondary"}>
            {childrenArray[0]}
            <div className={"absolute -top-3 font-medium text-base left-0 text-center w-full"}>
                {childrenArray[1]}
            </div>
        </div>
    )
}

export default LeftSignSideOrTextLayout;