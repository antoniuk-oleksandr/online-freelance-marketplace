import MyInput from "@/common-components/MyInput/MyInput";
import {getErrorMessage} from "@/utils/input-utils";
import {useFormContext} from "react-hook-form";

const SignUpInputs = () => {
    const {register, formState} = useFormContext();
    const {errors} = formState;

    return (
        <>
            <MyInput
                register={register}
                name={"name"}
                error={getErrorMessage(errors, 'name')}
                label={"Your name"}
            />
            <MyInput
                register={register}
                name={"email"}
                error={getErrorMessage(errors, 'email')}
                label={"Email"}
            />
            <MyInput
                inputType={"password"}
                register={register}
                name={"password"}
                error={getErrorMessage(errors, 'password')}
                label={"Password"}
            />
        </>
    )
}

export default SignUpInputs;