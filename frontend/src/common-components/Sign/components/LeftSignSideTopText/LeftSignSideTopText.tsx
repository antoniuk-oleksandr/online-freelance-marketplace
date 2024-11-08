import LeftSignSideTopTextLayout from "./LeftSignSideTopTextLayout";
import {SignProps} from "@/types/SignProps";

const LeftSignSideTopText = (props: SignProps) => {
    const {signText, subSignText} = props;

    return (
        <LeftSignSideTopTextLayout>
            <h2 className={"font-bold text-4xl "}>{signText}</h2>
            <p
                className={"text-lg text-light-palette-text-secondary dark:text-dark-palette-text-secondary"}
            >{subSignText}</p>
        </LeftSignSideTopTextLayout>
    )
}

export default LeftSignSideTopText;