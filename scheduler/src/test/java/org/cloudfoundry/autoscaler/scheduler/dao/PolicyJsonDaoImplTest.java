package org.cloudfoundry.autoscaler.scheduler.dao;

import static org.hamcrest.Matchers.is;
import static org.junit.Assert.assertThat;

import java.io.IOException;
import java.util.List;
import java.util.UUID;

import org.cloudfoundry.autoscaler.scheduler.entity.PolicyJsonEntity;
import org.cloudfoundry.autoscaler.scheduler.util.ConsulUtil;
import org.cloudfoundry.autoscaler.scheduler.util.TestDataDbUtil;
import org.junit.AfterClass;
import org.junit.Before;
import org.junit.BeforeClass;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

@RunWith(SpringRunner.class)
@SpringBootTest

public class PolicyJsonDaoImplTest{

	@Autowired
	private PolicyJsonDao policyJsonDao;
	
	@Autowired
	private TestDataDbUtil testDataDbUtil;
	
	private static ConsulUtil consulUtil;
	
	@BeforeClass
	public static void beforeClass() throws IOException{
		consulUtil = new ConsulUtil();
		consulUtil.start();
		
	}
	@AfterClass
	public static void afterClass() throws IOException, InterruptedException{
		consulUtil.stop();
	}
	@Before
	public void before() throws IOException{
		String appId = "the_app_id";
		String guid = UUID.randomUUID().toString();
		testDataDbUtil.cleanupData();
		testDataDbUtil.insertPolicyJson(appId, guid);
	}
	@Test
	public void testGetAllPolicies(){
		List<PolicyJsonEntity> policyJsonList = policyJsonDao.getAllPolicies();
		assertThat("It should have 1 policy_json", policyJsonList.size(),is(1));
		PolicyJsonEntity entity = policyJsonList.get(0);
		assertThat("The app_id of entity should equal", entity.getAppId(), is("the_app_id"));
	}
}
