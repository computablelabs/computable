const CONSTANTS = require("./migration_constants");

module.exports = (deployer, network, accounts) => {
    const OWNER = accounts[0];

    // Import the Dharma contracts.
    const AttributeStore = artifacts.require("AttributeStore");
    const BasicToken = artifacts.require("BasicToken");
    const ConstructableToken = artifacts.require("ConstructableToken");
    const DLL = artifacts.require("DLL");
    const ERC20 = artifacts.require("ERC20");
    const ERC20Basic = artifacts.require("ERC20Basic");
    const PLCRVoting = artifacts.require("PLCRVoting");
    const Parameterizer = artifacts.require("Paramterizer");
    const Registry = artifacts.require("Registry");
    const SafeMath = artifacts.require("SafeMath");
    const StandardToken = artifacts.require("StandardToken");

    //const {
    //    signatories,
    //    numAuthorizationsRequired,
    //    timelock,
    //} = generateParamsForDharmaMultiSigWallet(network, accounts);

    //// Deploy the DharmaMultiSigWallet with a set of signatories, the number of
    //// authorizations required before a transaction can be executed, and the
    //// timelock period, defined in seconds.
    //deployer.deploy(DharmaMultiSigWallet, signatories, numAuthorizationsRequired, timelock);

    //// Deploy our Permissions library and link it to the contracts in our protocol that depend on it.
    //deployer.deploy(PermissionsLib);
    //deployer.link(PermissionsLib, [DebtRegistry, TokenTransferProxy, Collateralizer, DebtToken]);
    deployer.deploy(SafeMath);
    deployer.deploy(DLL);

    return deployer.deploy(DebtRegistry).then(async () => {
        await deployer.deploy(DebtToken, DebtRegistry.address);
        await deployer.deploy(TokenTransferProxy);
        await deployer.deploy(RepaymentRouter, DebtRegistry.address, TokenTransferProxy.address);
        await deployer.deploy(DebtKernel, TokenTransferProxy.address);
        await deployer.deploy(TokenRegistry).then(async () => {
            const DummyToken = artifacts.require("DummyToken");
            await configureTokenRegistry(network, accounts, TokenRegistry, DummyToken);
        });
        await deployer.deploy(
            Collateralizer,
            DebtKernel.address,
            DebtRegistry.address,
            TokenRegistry.address,
            TokenTransferProxy.address,
        );
        await deployer.deploy(
            ContractRegistry,
            Collateralizer.address,
            DebtKernel.address,
            DebtRegistry.address,
            DebtToken.address,
            RepaymentRouter.address,
            TokenRegistry.address,
            TokenTransferProxy.address,
        );
    });
};
