// Specifically request an abstraction for DataCoin
var DataCoin = artifacts.require('DataCoin');

contract('DataCoin', function (accounts) {
  it('should put 10000 DataCoin in the first account', function () {
    return DataCoin.deployed(accounts[0]).then(function (instance) {
      return instance.balanceOf.call(accounts[0]);
    }).then(function (balance) {
      assert.equal(balance.valueOf(), 10000, '10000 wasn\'t in the first account');
    });
  });
  it('should send coin correctly', function () {
    var data;

    // Get initial balances of first and second account.
    var accountOne = accounts[0];
    var accountTwo = accounts[1];

    var accountOneStartingBalance;
    var accountTwoStartingBalance;
    var accountOneEndingBalance;
    var accountTwoEndingBalance;

    var amount = 10;

    return DataCoin.deployed().then(function (instance) {
      data = instance;
      return data.balanceOf.call(accountOne);
    }).then(function (balance) {
      accountOneStartingBalance = balance.toNumber();
      return data.balanceOf.call(accountTwo);
    }).then(function (balance) {
      accountTwoStartingBalance = balance.toNumber();
      return data.transfer(accountTwo, amount, { from: accountOne });
    }).then(function () {
      return data.balanceOf.call(accountOne);
    }).then(function (balance) {
      accountOneEndingBalance = balance.toNumber();
      return data.balanceOf.call(accountTwo);
    }).then(function (balance) {
      accountTwoEndingBalance = balance.toNumber();

      assert.equal(accountOneEndingBalance,
        accountOneStartingBalance - amount,
        'Amount wasn\'t correctly taken from the sender');
      assert.equal(accountTwoEndingBalance,
        accountTwoStartingBalance + amount,
        'Amount wasn\'t correctly sent to the receiver');
    });
  });
});
