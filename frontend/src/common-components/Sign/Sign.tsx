import SignLayout from "./SignLayout";
import LeftSignSide from "@/common-components/Sign/components/LeftSignSide/LeftSignSide";
import RightSignSide from "@/common-components/Sign/components/RightSignSide/RightSignSide";
import {SignProps} from "@/types/SignProps";

const Sign = (props: SignProps) => {
    return (
        <SignLayout>
            <LeftSignSide {...props}/>
            <RightSignSide/>
        </SignLayout>
    )
}

export default Sign;