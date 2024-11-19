import FooterIconLayout from "./FooterIconLayout";
import {ReactNode} from "react";

type FooterIconProps = {
    icon: ReactNode,
    link: string,
}

const FooterIcon = (props: FooterIconProps) => {
    const {icon} = props;

    return (
        <FooterIconLayout {...props}>
            {icon}
        </FooterIconLayout>
    )
}

export default FooterIcon;