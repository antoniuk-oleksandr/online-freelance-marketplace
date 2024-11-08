import {SignProps} from "@/types/SignProps";
import {handleGoogleSignInButtonClick} from "@/common-components/Sign/handlers";
import {useGoogleLogin} from "@react-oauth/google";

const GoogleButton = (props: SignProps) => {
    const {googleButtonText} = props;

    const onClick = useGoogleLogin({
        onSuccess: tokenResponse => handleGoogleSignInButtonClick(tokenResponse),
    });

    return (
        <button
            onClick={() => onClick()}
            type={"button"}
            className={"flex !h-11 font-semibold hover:bg-light-palette-action-hover dark:hover:bg-dark-palette-action-hover duration-200 ease-out active:scale-95 items-center justify-center border border-light-palette-divider dark:border-dark-palette-divider rounded-sm  gap-x-2 bg-light-palette-background-block dark:bg-dark-palette-background-block w-full"}
        >
            <span>{googleButtonText}</span>
            <img src={"/images/plus.png"} alt={"google"}/>
        </button>
    )
}

export default GoogleButton;