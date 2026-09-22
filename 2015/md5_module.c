#include <janet.h>
#include <CommonCrypto/CommonDigest.h>

static Janet cfun_md5(int32_t argc, Janet *argv) {
    janet_fixarity(argc, 1);
    const uint8_t *input = janet_getstring(argv, 0);
    int32_t len = janet_string_length(input);

    unsigned char digest[CC_MD5_DIGEST_LENGTH];
    CC_MD5_CTX ctx;
    CC_MD5_Init(&ctx);
    CC_MD5_Update(&ctx, input, (CC_LONG)len);
    CC_MD5_Final(digest, &ctx);

    char hex[33];
    for (int i = 0; i < CC_MD5_DIGEST_LENGTH; i++)
        sprintf(hex + i * 2, "%02x", digest[i]);

    return janet_cstringv(hex);
}

static const JanetReg cfuns[] = {
    {"md5", cfun_md5, "(md5 str)\n\nCompute the MD5 hex digest of str."},
    {NULL, NULL, NULL}
};

JANET_MODULE_ENTRY(JanetTable *env) {
    janet_cfuns(env, "md5", cfuns);
}
