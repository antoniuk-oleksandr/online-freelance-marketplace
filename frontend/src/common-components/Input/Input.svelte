<script lang="ts">
    import InputLayout from "@/common-components/Input/InputLayout.svelte";
    import InputError from "@/common-components/Input/components/InputError/InputError.svelte";
    import InputLabel from "@/common-components/Input/components/InputLabel/InputLabel.svelte";
    import InputElement from "@/common-components/Input/components/InputElement/InputElement.svelte";
    import InputContainerLayout from "@/common-components/Input/components/InputContainer/InputContainerLayout.svelte";

    type InputProps = {
        id: string,
        type: "text" | "password",
        label: string,
        value: string,
        error: string | null,
        wasSubmitted: boolean,
    }

    let {id, type, label, value, error, wasSubmitted}: InputProps = $props();

    let isFocused = $state(false);
    const setIsFocused = (value: boolean) => isFocused = value;
</script>

<InputLayout>
    <InputContainerLayout error={error} isFocused={isFocused} wasSubmitted={wasSubmitted}>
        <InputLabel
                id={id} value={value}
                label={label} error={error}
                wasSubmitted={wasSubmitted}
                isFocused={isFocused}
        />
        <InputElement setIsFocused={setIsFocused} id={id} type={type}/>
    </InputContainerLayout>
    <InputError wasSubmitted={wasSubmitted} error={error}/>
</InputLayout>

