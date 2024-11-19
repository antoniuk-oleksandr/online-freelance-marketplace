import {MantineSize, PasswordInput, PasswordInputProps, TextInput, TextInputProps} from "@mantine/core";
import {InputHTMLAttributes} from "react";
import {RegisterOptions, UseFormRegister} from "react-hook-form";

type MyInputProps = {
    size?: MantineSize,
    register?: UseFormRegister<any>,
    rules?: RegisterOptions,
    inputType?: 'password' | 'text',
} & InputHTMLAttributes<HTMLInputElement> & TextInputProps & PasswordInputProps;

const MyInput = (props: MyInputProps) => {
    const {register, name, rules, inputType, ...rest} = props;
    const inputProps = register && name ? register(name, rules) : {};

    if (inputType === 'text' || !inputType) return (
        <TextInput
            {...inputProps}
            {...rest}
            autoComplete={"off"}
        />
    )
    else if(inputType === 'password') return (
        <PasswordInput
            {...inputProps}
            {...rest}
            autoComplete={"off"}
        />
    )
}

export default MyInput;