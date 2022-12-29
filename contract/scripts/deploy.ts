const main = async () => {
    //@ts-ignore hre is injected by hardhat runtime environment
    const [deployer] = await hre.ethers.getSigners();
    const accountBalance = await deployer.getBalance();
  
    console.log("Deploying contracts with account: ", deployer.address);
    console.log("Account balance: ", accountBalance.toString());
  
    //@ts-ignore hre is injected by hardhat runtime environment
    const electionsContractFactory = await hre.ethers.getContractFactory("Elections");
    const electionContract = await electionsContractFactory.deploy();
    await electionContract.deployed();
  
    console.log("WavePortal address: ", electionContract.address);
  };
  
  const runMain = async () => {
    try {
      await main();
      process.exit(0);
    } catch (error) {
      console.log(error);
      process.exit(1);
    }
  };
  
  runMain();