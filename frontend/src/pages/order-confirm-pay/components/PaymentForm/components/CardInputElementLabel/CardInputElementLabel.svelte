<script lang="ts">
    type CardInputElementLabelProps = {
        label: string,
        id: string,
    }

    const {label, id}: CardInputElementLabelProps = $props();
</script>

<label for={id}>{label}</label>
