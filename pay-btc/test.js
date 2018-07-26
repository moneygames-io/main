const bcoin = require('bcoin');
const Keyring = bcoin.wallet.WalletKey;

const master = bcoin.hd.generate();

const key = master.derivePath('m/44/0/0/0/0');

