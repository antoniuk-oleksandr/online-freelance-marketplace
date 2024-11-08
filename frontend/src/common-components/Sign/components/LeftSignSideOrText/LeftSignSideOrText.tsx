import LeftSignSideOrTextLayout from "./LeftSignSideOrTextLayout";
import DividerElement from "@/common-components/DividerElement";

const LeftSignSideOrText = () => {
    return (
        <LeftSignSideOrTextLayout>
            <DividerElement/>
            <span className={"bg-light-palette-background-block dark:bg-dark-palette-background-block px-2.5"}>or</span>
        </LeftSignSideOrTextLayout>
    )
}

export default LeftSignSideOrText;