package main.java.org.app.chaincode.invocation;

import static java.nio.charset.StandardCharsets.UTF_8;

import java.util.Collection;
import java.util.logging.Level;
import java.util.logging.Logger;

import org.hyperledger.fabric.sdk.Channel;
import org.hyperledger.fabric.sdk.EventHub;
import org.hyperledger.fabric.sdk.Orderer;
import org.hyperledger.fabric.sdk.Peer;
import org.hyperledger.fabric.sdk.ProposalResponse;

import main.java.org.app.client.CAClient;
import main.java.org.app.client.ChannelClient;
import main.java.org.app.client.FabricClient;
import main.java.org.app.config.Config;
import main.java.org.app.user.UserContext;
import main.java.org.app.util.Util;

public class QueryDoctors {
	private static final byte[] EXPECTED_EVENT_DATA = "!".getBytes(UTF_8);
	private static final String EXPECTED_EVENT_NAME = "event";

	public static void main(String args[]) {
		try {
            Util.cleanUp();
			String caUrl = Config.CA_ORG1_URL;
			CAClient caClient = new CAClient(caUrl, null);
			// Enroll Admin to Org1MSP
			UserContext adminUserContext = new UserContext();
			adminUserContext.setName(Config.ADMIN);
			adminUserContext.setAffiliation(Config.ORG1);
			adminUserContext.setMspId(Config.ORG1_MSP);
			caClient.setAdminUserContext(adminUserContext);
			adminUserContext = caClient.enrollAdminUser(Config.ADMIN, Config.ADMIN_PASSWORD);
			
			FabricClient fabClient = new FabricClient(adminUserContext);
			
			ChannelClient channelClient = fabClient.createChannelClient(Config.CHANNEL_NAME);
			Channel channel = channelClient.getChannel();
			Peer peer = fabClient.getInstance().newPeer(Config.ORG1_PEER_0, Config.ORG1_PEER_0_URL);
			EventHub eventHub = fabClient.getInstance().newEventHub("eventhub01", "grpc://localhost:7053");
			Orderer orderer = fabClient.getInstance().newOrderer(Config.ORDERER_NAME, Config.ORDERER_URL);
			channel.addPeer(peer);
			channel.addEventHub(eventHub);
			channel.addOrderer(orderer);
			channel.initialize();

			Logger.getLogger(QueryPatient.class.getName()).log(Level.INFO, "Querying for all Doctors ...");
			Collection<ProposalResponse>  responsesQuery = channelClient.queryByChainCode(Config.CHAINCODE_1_NAME, "queryAllDoctors", null);
			for (ProposalResponse pres : responsesQuery) {
				String stringResponse = new String(pres.getChaincodeActionResponsePayload());
				Logger.getLogger(QueryPatient.class.getName()).log(Level.INFO, stringResponse);
			}

			Thread.sleep(1000);
			String[] args1 = {"Waqas"};
			Logger.getLogger(QueryPatient.class.getName()).log(Level.INFO, "Querying for a Doctor - " + args1[0]);
			
			Collection<ProposalResponse>  responses1Query = channelClient.queryByChainCode(Config.CHAINCODE_1_NAME, "queryDoctor", args1);
			for (ProposalResponse pres : responses1Query) {
				String stringResponse = new String(pres.getChaincodeActionResponsePayload());
				Logger.getLogger(QueryPatient.class.getName()).log(Level.INFO, stringResponse);
			}		
			
		} catch (Exception e) {
			e.printStackTrace();
		}
	}
}
