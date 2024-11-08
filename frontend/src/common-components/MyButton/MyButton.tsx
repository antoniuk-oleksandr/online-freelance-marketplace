import {Button, ButtonProps} from "@mantine/core";
import {LayoutProps} from "@/types/LayoutProps";
import {ButtonHTMLAttributes} from "react";

type MyButtonProps = {
} & LayoutProps & ButtonProps & ButtonHTMLAttributes<HTMLButtonElement>;

const MyButton = (props: MyButtonProps) => {
    const {children, className, ...rest} = props;

    return (
        <Button
            {...rest}
            className={`${className} !h-12 !grid !place-items-center !bg-cyan-500 hover:!bg-cyan-400 active:!scale-95 !duration-200 font-semibold !ease-out`}
        >
            {children}
        </Button>
    )
}

export default MyButton;