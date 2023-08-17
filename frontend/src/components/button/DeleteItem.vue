<template>
    <button class="noselect"><span class="text">Delete</span><span class="icon"><svg xmlns="http://www.w3.org/2000/svg"
                width="24" height="24" viewBox="0 0 24 24">
                <path
                    d="M24 20.188l-8.315-8.209 8.2-8.282-3.697-3.697-8.212 8.318-8.31-8.203-3.666 3.666 8.321 8.24-8.206 8.313 3.666 3.666 8.237-8.318 8.285 8.203z">
                </path>
            </svg></span></button>
</template>

<script lang="ts" setup>
/* script here */
</script>

<style lang="less" scoped>
@var1: #e62222;
@var2: #c41b1b;
@var3: #eeeeee;
@var4: #ff3636;

button {
    width: 120px;
    height: 30px;
    cursor: pointer;
    display: flex;
    align-items: center;
    background: red;
    border: none;
    border-radius: 5px;
    box-shadow: 1px 1px 3px rgba(0, 0, 0, 0.15);
    background: @var1;

    .text {
        transform: translateX(25px);
        color: white;
        font-weight: bold;
    }

    .icon {
        position: absolute;
        border-left: 1px solid @var2;
        transform: translateX(90px);
        // height: 40px;
        width: 30px;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    svg {
        width: 15px;
        fill: @var3;
    }

    &:hover {
        background: @var4;

        .text {
            color: transparent;
        }

        .icon {
            width: 120px;
            border-left: none;
            transform: translateX(0);
        }
    }

    &:focus {
        outline: none;
    }

    &:active .icon svg {
        transform: scale(0.8);
    }
}

button,
button span {
    transition: 200ms;
}
</style>