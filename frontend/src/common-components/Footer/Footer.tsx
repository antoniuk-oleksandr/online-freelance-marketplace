import FooterLayout from "./FooterLayout";
import AppLogo from "@/common-components/AppLogo";
import {FaFacebook, FaInstagram, FaLinkedinIn, FaXTwitter} from "react-icons/fa6";
import FooterIcon from "@/common-components/Footer/components/FooterIcon/FooterIcon";

const Footer = () => {
    return (
        <FooterLayout>
            <div className={"flex gap-x-6 items-center flex-col sm:flex-row gap-y-4"}>
                <AppLogo/>
                <p>© 2024 Online Freelance Marketplace. All rights reserved.</p>
            </div>
            <div className={"text-xl flex gap-x-6"}>
                <FooterIcon icon={<FaInstagram/>} link={"https://www.instagram.com/"}/>
                <FooterIcon icon={<FaFacebook/>} link={"https://www.facebook.com/"}/>
                <FooterIcon icon={<FaLinkedinIn/>} link={"https://www.linkedin.com/"}/>
                <FooterIcon icon={<FaInstagram/>} link={"https://www.instagram.com/"}/>
                <FooterIcon icon={<FaXTwitter/>} link={"https://x.com/"}/>
            </div>
        </FooterLayout>
    )
}

export default Footer;