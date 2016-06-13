package ec2instances

const data = `
[
  {
    "family": "General Purpose",
    "enhanced_networking": false,
    "vCPU": 1,
    "generation": "previous",
    "ebs_iops": 0,
    "network_performance": "Low",
    "ebs_throughput": 0,
    "pretty_name": "M1 General Purpose Small",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.094",
            "yrTerm3.allUpfront": "0.0623",
            "yrTerm1.allUpfront": "0.0783",
            "yrTerm1.partialUpfront": "0.08",
            "yrTerm3.partialUpfront": "0.0663"
          },
          "ondemand": "0.139"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.473",
            "yrTerm3.allUpfront": "0.3666",
            "yrTerm1.allUpfront": "0.396",
            "yrTerm1.partialUpfront": "0.4041",
            "yrTerm3.partialUpfront": "0.39"
          },
          "ondemand": "0.629"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.057",
            "yrTerm3.allUpfront": "0.0365",
            "yrTerm1.allUpfront": "0.0478",
            "yrTerm1.partialUpfront": "0.0488",
            "yrTerm3.partialUpfront": "0.0389"
          },
          "ondemand": "0.075"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.028",
            "yrTerm3.allUpfront": "0.0152",
            "yrTerm1.allUpfront": "0.0235",
            "yrTerm1.partialUpfront": "0.024",
            "yrTerm3.partialUpfront": "0.0162"
          },
          "ondemand": "0.044"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.108",
            "yrTerm3.allUpfront": "0.0731",
            "yrTerm1.allUpfront": "0.0902",
            "yrTerm1.partialUpfront": "0.092",
            "yrTerm3.partialUpfront": "0.0778"
          },
          "ondemand": "0.139"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.525",
            "yrTerm3.allUpfront": "0.406",
            "yrTerm1.allUpfront": "0.4394",
            "yrTerm1.partialUpfront": "0.4483",
            "yrTerm3.partialUpfront": "0.4319"
          },
          "ondemand": "0.672"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.066",
            "yrTerm3.allUpfront": "0.043",
            "yrTerm1.allUpfront": "0.0551",
            "yrTerm1.partialUpfront": "0.0563",
            "yrTerm3.partialUpfront": "0.0458"
          },
          "ondemand": "0.088"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.037",
            "yrTerm3.allUpfront": "0.0205",
            "yrTerm1.allUpfront": "0.0317",
            "yrTerm1.partialUpfront": "0.0324",
            "yrTerm3.partialUpfront": "0.0219"
          },
          "ondemand": "0.061"
        }
      },
      "sa-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.134",
            "yrTerm3.allUpfront": "0.0798",
            "yrTerm1.allUpfront": "0.1123",
            "yrTerm1.partialUpfront": "0.1147",
            "yrTerm3.partialUpfront": "0.0849"
          },
          "ondemand": "0.168"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.513",
            "yrTerm3.allUpfront": "0.3841",
            "yrTerm1.allUpfront": "0.43",
            "yrTerm1.partialUpfront": "0.4388",
            "yrTerm3.partialUpfront": "0.4086"
          },
          "ondemand": "0.664"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.082",
            "yrTerm3.allUpfront": "0.0464",
            "yrTerm1.allUpfront": "0.0776",
            "yrTerm1.partialUpfront": "0.0792",
            "yrTerm3.partialUpfront": "0.0494"
          },
          "ondemand": "0.089"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.034",
            "yrTerm3.allUpfront": "0.0203",
            "yrTerm1.allUpfront": "0.0287",
            "yrTerm1.partialUpfront": "0.0294",
            "yrTerm3.partialUpfront": "0.0216"
          },
          "ondemand": "0.058"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.105",
            "yrTerm3.allUpfront": "0.0689",
            "yrTerm1.allUpfront": "0.0876",
            "yrTerm1.partialUpfront": "0.0894",
            "yrTerm3.partialUpfront": "0.0733"
          },
          "ondemand": "0.139"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.484",
            "yrTerm3.allUpfront": "0.3732",
            "yrTerm1.allUpfront": "0.4053",
            "yrTerm1.partialUpfront": "0.4135",
            "yrTerm3.partialUpfront": "0.3971"
          },
          "ondemand": "0.629"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.065",
            "yrTerm3.allUpfront": "0.0412",
            "yrTerm1.allUpfront": "0.0547",
            "yrTerm1.partialUpfront": "0.0558",
            "yrTerm3.partialUpfront": "0.0439"
          },
          "ondemand": "0.093"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.039",
            "yrTerm3.allUpfront": "0.0207",
            "yrTerm1.allUpfront": "0.0326",
            "yrTerm1.partialUpfront": "0.0334",
            "yrTerm3.partialUpfront": "0.0221"
          },
          "ondemand": "0.058"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.105",
            "yrTerm3.allUpfront": "0.0689",
            "yrTerm1.allUpfront": "0.0876",
            "yrTerm1.partialUpfront": "0.0894",
            "yrTerm3.partialUpfront": "0.0733"
          },
          "ondemand": "0.139"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.484",
            "yrTerm3.allUpfront": "0.3732",
            "yrTerm1.allUpfront": "0.4053",
            "yrTerm1.partialUpfront": "0.4135",
            "yrTerm3.partialUpfront": "0.3971"
          },
          "ondemand": "0.629"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.065",
            "yrTerm3.allUpfront": "0.0412",
            "yrTerm1.allUpfront": "0.0547",
            "yrTerm1.partialUpfront": "0.0558",
            "yrTerm3.partialUpfront": "0.0439"
          },
          "ondemand": "0.093"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.039",
            "yrTerm3.allUpfront": "0.0207",
            "yrTerm1.allUpfront": "0.0326",
            "yrTerm1.partialUpfront": "0.0334",
            "yrTerm3.partialUpfront": "0.0221"
          },
          "ondemand": "0.058"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.094",
            "yrTerm3.allUpfront": "0.0623",
            "yrTerm1.allUpfront": "0.0783",
            "yrTerm1.partialUpfront": "0.08",
            "yrTerm3.partialUpfront": "0.0663"
          },
          "ondemand": "0.139"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.473",
            "yrTerm3.allUpfront": "0.3666",
            "yrTerm1.allUpfront": "0.396",
            "yrTerm1.partialUpfront": "0.4041",
            "yrTerm3.partialUpfront": "0.39"
          },
          "ondemand": "0.629"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.057",
            "yrTerm3.allUpfront": "0.0365",
            "yrTerm1.allUpfront": "0.0478",
            "yrTerm1.partialUpfront": "0.0488",
            "yrTerm3.partialUpfront": "0.0389"
          },
          "ondemand": "0.075"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.028",
            "yrTerm3.allUpfront": "0.0152",
            "yrTerm1.allUpfront": "0.0235",
            "yrTerm1.partialUpfront": "0.024",
            "yrTerm3.partialUpfront": "0.0162"
          },
          "ondemand": "0.044"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.091",
            "yrTerm3.allUpfront": "0.0592",
            "yrTerm1.allUpfront": "0.0703",
            "yrTerm1.partialUpfront": "0.0774",
            "yrTerm3.partialUpfront": "0.0651"
          },
          "ondemand": "0.159"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.481",
            "yrTerm3.allUpfront": "0.361",
            "yrTerm1.allUpfront": "0.3743",
            "yrTerm1.partialUpfront": "0.4114",
            "yrTerm3.partialUpfront": "0.3967"
          },
          "ondemand": "0.653"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.064",
            "yrTerm3.allUpfront": "0.0387",
            "yrTerm1.allUpfront": "0.0494",
            "yrTerm1.partialUpfront": "0.0548",
            "yrTerm3.partialUpfront": "0.0438"
          },
          "ondemand": "0.084"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.033",
            "yrTerm3.allUpfront": "0.0176",
            "yrTerm1.allUpfront": "0.0264",
            "yrTerm1.partialUpfront": "0.0282",
            "yrTerm3.partialUpfront": "0.0193"
          },
          "ondemand": "0.053"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.105",
            "yrTerm3.allUpfront": "0.0689",
            "yrTerm1.allUpfront": "0.0876",
            "yrTerm1.partialUpfront": "0.0894",
            "yrTerm3.partialUpfront": "0.0733"
          },
          "ondemand": "0.147"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.484",
            "yrTerm3.allUpfront": "0.3732",
            "yrTerm1.allUpfront": "0.4053",
            "yrTerm1.partialUpfront": "0.4135",
            "yrTerm3.partialUpfront": "0.3971"
          },
          "ondemand": "0.639"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.066",
            "yrTerm3.allUpfront": "0.0412",
            "yrTerm1.allUpfront": "0.0557",
            "yrTerm1.partialUpfront": "0.0568",
            "yrTerm3.partialUpfront": "0.0439"
          },
          "ondemand": "0.078"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.034",
            "yrTerm3.allUpfront": "0.0189",
            "yrTerm1.allUpfront": "0.0293",
            "yrTerm1.partialUpfront": "0.03",
            "yrTerm3.partialUpfront": "0.0202"
          },
          "ondemand": "0.047"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.094",
            "yrTerm3.allUpfront": "0.0623",
            "yrTerm1.allUpfront": "0.0783",
            "yrTerm1.partialUpfront": "0.08",
            "yrTerm3.partialUpfront": "0.0663"
          },
          "ondemand": "0.139"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.473",
            "yrTerm3.allUpfront": "0.3666",
            "yrTerm1.allUpfront": "0.396",
            "yrTerm1.partialUpfront": "0.4041",
            "yrTerm3.partialUpfront": "0.39"
          },
          "ondemand": "0.629"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.066",
            "yrTerm3.allUpfront": "0.0412",
            "yrTerm1.allUpfront": "0.0557",
            "yrTerm1.partialUpfront": "0.0568",
            "yrTerm3.partialUpfront": "0.0439"
          },
          "ondemand": "0.075"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.034",
            "yrTerm3.allUpfront": "0.0189",
            "yrTerm1.allUpfront": "0.0293",
            "yrTerm1.partialUpfront": "0.03",
            "yrTerm3.partialUpfront": "0.0202"
          },
          "ondemand": "0.047"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 4,
      "max_enis": 2
    },
    "arch": [
      "i386",
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "PV"
    ],
    "ebs_optimized": false,
    "storage": {
      "ssd": false,
      "devices": 1,
      "size": 160.0
    },
    "instance_type": "m1.small",
    "ECU": 1.0,
    "memory": 1.7,
    "ebs_max_bandwidth": 0
  },
  {
    "family": "General purpose",
    "enhanced_networking": false,
    "vCPU": 1,
    "generation": "previous",
    "ebs_iops": 0,
    "network_performance": "Moderate",
    "ebs_throughput": 0,
    "pretty_name": "M1 General Purpose Medium",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.155",
            "yrTerm3.allUpfront": "0.0986",
            "yrTerm1.allUpfront": "0.1296",
            "yrTerm1.partialUpfront": "0.1322",
            "yrTerm3.partialUpfront": "0.1049"
          },
          "ondemand": "0.237"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.543",
            "yrTerm3.allUpfront": "0.4088",
            "yrTerm1.allUpfront": "0.4547",
            "yrTerm1.partialUpfront": "0.4639",
            "yrTerm3.partialUpfront": "0.4349"
          },
          "ondemand": "0.744"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.115",
            "yrTerm3.allUpfront": "0.0721",
            "yrTerm1.allUpfront": "0.0963",
            "yrTerm1.partialUpfront": "0.0983",
            "yrTerm3.partialUpfront": "0.0767"
          },
          "ondemand": "0.149"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.056",
            "yrTerm3.allUpfront": "0.0294",
            "yrTerm1.allUpfront": "0.0471",
            "yrTerm1.partialUpfront": "0.0482",
            "yrTerm3.partialUpfront": "0.0313"
          },
          "ondemand": "0.087"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.198",
            "yrTerm3.allUpfront": "0.1162",
            "yrTerm1.allUpfront": "0.1658",
            "yrTerm1.partialUpfront": "0.1692",
            "yrTerm3.partialUpfront": "0.1236"
          },
          "ondemand": "0.237"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.625",
            "yrTerm3.allUpfront": "0.4561",
            "yrTerm1.allUpfront": "0.5234",
            "yrTerm1.partialUpfront": "0.5341",
            "yrTerm3.partialUpfront": "0.4853"
          },
          "ondemand": "0.782"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.133",
            "yrTerm3.allUpfront": "0.0849",
            "yrTerm1.allUpfront": "0.1116",
            "yrTerm1.partialUpfront": "0.1139",
            "yrTerm3.partialUpfront": "0.0904"
          },
          "ondemand": "0.177"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.074",
            "yrTerm3.allUpfront": "0.0411",
            "yrTerm1.allUpfront": "0.0645",
            "yrTerm1.partialUpfront": "0.0658",
            "yrTerm3.partialUpfront": "0.0437"
          },
          "ondemand": "0.122"
        }
      },
      "sa-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.214",
            "yrTerm3.allUpfront": "0.1263",
            "yrTerm1.allUpfront": "0.1793",
            "yrTerm1.partialUpfront": "0.1831",
            "yrTerm3.partialUpfront": "0.1344"
          },
          "ondemand": "0.297"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.615",
            "yrTerm3.allUpfront": "0.4447",
            "yrTerm1.allUpfront": "0.5155",
            "yrTerm1.partialUpfront": "0.526",
            "yrTerm3.partialUpfront": "0.4731"
          },
          "ondemand": "0.814"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.165",
            "yrTerm3.allUpfront": "0.0932",
            "yrTerm1.allUpfront": "0.1549",
            "yrTerm1.partialUpfront": "0.1581",
            "yrTerm3.partialUpfront": "0.0992"
          },
          "ondemand": "0.179"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.068",
            "yrTerm3.allUpfront": "0.0397",
            "yrTerm1.allUpfront": "0.0565",
            "yrTerm1.partialUpfront": "0.0577",
            "yrTerm3.partialUpfront": "0.0423"
          },
          "ondemand": "0.117"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.173",
            "yrTerm3.allUpfront": "0.1113",
            "yrTerm1.allUpfront": "0.1447",
            "yrTerm1.partialUpfront": "0.1477",
            "yrTerm3.partialUpfront": "0.1184"
          },
          "ondemand": "0.237"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.563",
            "yrTerm3.allUpfront": "0.4217",
            "yrTerm1.allUpfront": "0.4711",
            "yrTerm1.partialUpfront": "0.4808",
            "yrTerm3.partialUpfront": "0.4486"
          },
          "ondemand": "0.744"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.13",
            "yrTerm3.allUpfront": "0.0825",
            "yrTerm1.allUpfront": "0.1091",
            "yrTerm1.partialUpfront": "0.1113",
            "yrTerm3.partialUpfront": "0.0877"
          },
          "ondemand": "0.187"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.078",
            "yrTerm3.allUpfront": "0.0414",
            "yrTerm1.allUpfront": "0.0645",
            "yrTerm1.partialUpfront": "0.0658",
            "yrTerm3.partialUpfront": "0.0441"
          },
          "ondemand": "0.117"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.173",
            "yrTerm3.allUpfront": "0.1113",
            "yrTerm1.allUpfront": "0.1447",
            "yrTerm1.partialUpfront": "0.1477",
            "yrTerm3.partialUpfront": "0.1184"
          },
          "ondemand": "0.237"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.563",
            "yrTerm3.allUpfront": "0.4217",
            "yrTerm1.allUpfront": "0.4711",
            "yrTerm1.partialUpfront": "0.4808",
            "yrTerm3.partialUpfront": "0.4486"
          },
          "ondemand": "0.744"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.13",
            "yrTerm3.allUpfront": "0.0825",
            "yrTerm1.allUpfront": "0.1091",
            "yrTerm1.partialUpfront": "0.1113",
            "yrTerm3.partialUpfront": "0.0877"
          },
          "ondemand": "0.187"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.078",
            "yrTerm3.allUpfront": "0.0414",
            "yrTerm1.allUpfront": "0.0645",
            "yrTerm1.partialUpfront": "0.0658",
            "yrTerm3.partialUpfront": "0.0441"
          },
          "ondemand": "0.117"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.155",
            "yrTerm3.allUpfront": "0.0986",
            "yrTerm1.allUpfront": "0.1296",
            "yrTerm1.partialUpfront": "0.1322",
            "yrTerm3.partialUpfront": "0.1049"
          },
          "ondemand": "0.237"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.543",
            "yrTerm3.allUpfront": "0.4088",
            "yrTerm1.allUpfront": "0.4547",
            "yrTerm1.partialUpfront": "0.4639",
            "yrTerm3.partialUpfront": "0.4349"
          },
          "ondemand": "0.744"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.115",
            "yrTerm3.allUpfront": "0.0721",
            "yrTerm1.allUpfront": "0.0963",
            "yrTerm1.partialUpfront": "0.0983",
            "yrTerm3.partialUpfront": "0.0767"
          },
          "ondemand": "0.149"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.056",
            "yrTerm3.allUpfront": "0.0294",
            "yrTerm1.allUpfront": "0.0471",
            "yrTerm1.partialUpfront": "0.0482",
            "yrTerm3.partialUpfront": "0.0313"
          },
          "ondemand": "0.087"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.211",
            "yrTerm3.allUpfront": "0.1226",
            "yrTerm1.allUpfront": "0.1644",
            "yrTerm1.partialUpfront": "0.1806",
            "yrTerm3.partialUpfront": "0.1347"
          },
          "ondemand": "0.279"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.615",
            "yrTerm3.allUpfront": "0.4305",
            "yrTerm1.allUpfront": "0.4787",
            "yrTerm1.partialUpfront": "0.526",
            "yrTerm3.partialUpfront": "0.4731"
          },
          "ondemand": "0.792"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.139",
            "yrTerm3.allUpfront": "0.081",
            "yrTerm1.allUpfront": "0.1067",
            "yrTerm1.partialUpfront": "0.1189",
            "yrTerm3.partialUpfront": "0.0918"
          },
          "ondemand": "0.168"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.066",
            "yrTerm3.allUpfront": "0.0352",
            "yrTerm1.allUpfront": "0.0518",
            "yrTerm1.partialUpfront": "0.0553",
            "yrTerm3.partialUpfront": "0.0385"
          },
          "ondemand": "0.106"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.173",
            "yrTerm3.allUpfront": "0.1113",
            "yrTerm1.allUpfront": "0.1447",
            "yrTerm1.partialUpfront": "0.1477",
            "yrTerm3.partialUpfront": "0.1184"
          },
          "ondemand": "0.255"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.563",
            "yrTerm3.allUpfront": "0.4217",
            "yrTerm1.allUpfront": "0.4711",
            "yrTerm1.partialUpfront": "0.4808",
            "yrTerm3.partialUpfront": "0.4486"
          },
          "ondemand": "0.764"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.131",
            "yrTerm3.allUpfront": "0.0825",
            "yrTerm1.allUpfront": "0.11",
            "yrTerm1.partialUpfront": "0.1123",
            "yrTerm3.partialUpfront": "0.0877"
          },
          "ondemand": "0.157"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.068",
            "yrTerm3.allUpfront": "0.0378",
            "yrTerm1.allUpfront": "0.0589",
            "yrTerm1.partialUpfront": "0.0602",
            "yrTerm3.partialUpfront": "0.0403"
          },
          "ondemand": "0.095"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.155",
            "yrTerm3.allUpfront": "0.0986",
            "yrTerm1.allUpfront": "0.1296",
            "yrTerm1.partialUpfront": "0.1322",
            "yrTerm3.partialUpfront": "0.1049"
          },
          "ondemand": "0.237"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.543",
            "yrTerm3.allUpfront": "0.4088",
            "yrTerm1.allUpfront": "0.4547",
            "yrTerm1.partialUpfront": "0.4639",
            "yrTerm3.partialUpfront": "0.4349"
          },
          "ondemand": "0.744"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.131",
            "yrTerm3.allUpfront": "0.0825",
            "yrTerm1.allUpfront": "0.11",
            "yrTerm1.partialUpfront": "0.1123",
            "yrTerm3.partialUpfront": "0.0877"
          },
          "ondemand": "0.149"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.068",
            "yrTerm3.allUpfront": "0.0378",
            "yrTerm1.allUpfront": "0.0589",
            "yrTerm1.partialUpfront": "0.0602",
            "yrTerm3.partialUpfront": "0.0403"
          },
          "ondemand": "0.095"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 6,
      "max_enis": 2
    },
    "arch": [
      "i386",
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "PV"
    ],
    "ebs_optimized": false,
    "storage": {
      "ssd": false,
      "devices": 1,
      "size": 410.0
    },
    "instance_type": "m1.medium",
    "ECU": 2.0,
    "memory": 3.75,
    "ebs_max_bandwidth": 0
  },
  {
    "family": "General purpose",
    "enhanced_networking": false,
    "vCPU": 2,
    "generation": "previous",
    "ebs_iops": 62.5,
    "network_performance": "Moderate",
    "ebs_throughput": 4000.0,
    "pretty_name": "M1 General Purpose Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.284",
            "yrTerm3.allUpfront": "0.1782",
            "yrTerm1.allUpfront": "0.2378",
            "yrTerm1.partialUpfront": "0.2426",
            "yrTerm3.partialUpfront": "0.1896"
          },
          "ondemand": "0.435"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.665",
            "yrTerm3.allUpfront": "0.4827",
            "yrTerm1.allUpfront": "0.5567",
            "yrTerm1.partialUpfront": "0.5681",
            "yrTerm3.partialUpfront": "0.5135"
          },
          "ondemand": "0.955"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.232",
            "yrTerm3.allUpfront": "0.1467",
            "yrTerm1.allUpfront": "0.1944",
            "yrTerm1.partialUpfront": "0.1983",
            "yrTerm3.partialUpfront": "0.156"
          },
          "ondemand": "0.299"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.112",
            "yrTerm3.allUpfront": "0.0588",
            "yrTerm1.allUpfront": "0.0953",
            "yrTerm1.partialUpfront": "0.0973",
            "yrTerm3.partialUpfront": "0.0625"
          },
          "ondemand": "0.175"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.332",
            "yrTerm3.allUpfront": "0.2117",
            "yrTerm1.allUpfront": "0.278",
            "yrTerm1.partialUpfront": "0.2836",
            "yrTerm3.partialUpfront": "0.2252"
          },
          "ondemand": "0.435"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.75",
            "yrTerm3.allUpfront": "0.5464",
            "yrTerm1.allUpfront": "0.6281",
            "yrTerm1.partialUpfront": "0.6408",
            "yrTerm3.partialUpfront": "0.5813"
          },
          "ondemand": "1.001"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.265",
            "yrTerm3.allUpfront": "0.1699",
            "yrTerm1.allUpfront": "0.2219",
            "yrTerm1.partialUpfront": "0.2264",
            "yrTerm3.partialUpfront": "0.1807"
          },
          "ondemand": "0.353"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.148",
            "yrTerm3.allUpfront": "0.0822",
            "yrTerm1.allUpfront": "0.1288",
            "yrTerm1.partialUpfront": "0.1314",
            "yrTerm3.partialUpfront": "0.0874"
          },
          "ondemand": "0.243"
        }
      },
      "sa-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.404",
            "yrTerm3.allUpfront": "0.232",
            "yrTerm1.allUpfront": "0.3385",
            "yrTerm1.partialUpfront": "0.3453",
            "yrTerm3.partialUpfront": "0.2468"
          },
          "ondemand": "0.553"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.826",
            "yrTerm3.allUpfront": "0.5611",
            "yrTerm1.allUpfront": "0.692",
            "yrTerm1.partialUpfront": "0.7061",
            "yrTerm3.partialUpfront": "0.597"
          },
          "ondemand": "1.092"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.331",
            "yrTerm3.allUpfront": "0.1844",
            "yrTerm1.allUpfront": "0.3118",
            "yrTerm1.partialUpfront": "0.3182",
            "yrTerm3.partialUpfront": "0.1962"
          },
          "ondemand": "0.357"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.136",
            "yrTerm3.allUpfront": "0.0805",
            "yrTerm1.allUpfront": "0.1131",
            "yrTerm1.partialUpfront": "0.1155",
            "yrTerm3.partialUpfront": "0.0856"
          },
          "ondemand": "0.233"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.322",
            "yrTerm3.allUpfront": "0.2025",
            "yrTerm1.allUpfront": "0.2701",
            "yrTerm1.partialUpfront": "0.2756",
            "yrTerm3.partialUpfront": "0.2154"
          },
          "ondemand": "0.435"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.706",
            "yrTerm3.allUpfront": "0.5079",
            "yrTerm1.allUpfront": "0.5916",
            "yrTerm1.partialUpfront": "0.6037",
            "yrTerm3.partialUpfront": "0.5403"
          },
          "ondemand": "0.955"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.261",
            "yrTerm3.allUpfront": "0.1636",
            "yrTerm1.allUpfront": "0.2188",
            "yrTerm1.partialUpfront": "0.2233",
            "yrTerm3.partialUpfront": "0.174"
          },
          "ondemand": "0.373"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.156",
            "yrTerm3.allUpfront": "0.082",
            "yrTerm1.allUpfront": "0.1279",
            "yrTerm1.partialUpfront": "0.1305",
            "yrTerm3.partialUpfront": "0.0872"
          },
          "ondemand": "0.233"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.322",
            "yrTerm3.allUpfront": "0.2025",
            "yrTerm1.allUpfront": "0.2701",
            "yrTerm1.partialUpfront": "0.2756",
            "yrTerm3.partialUpfront": "0.2154"
          },
          "ondemand": "0.435"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.706",
            "yrTerm3.allUpfront": "0.5079",
            "yrTerm1.allUpfront": "0.5916",
            "yrTerm1.partialUpfront": "0.6037",
            "yrTerm3.partialUpfront": "0.5403"
          },
          "ondemand": "0.955"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.261",
            "yrTerm3.allUpfront": "0.1636",
            "yrTerm1.allUpfront": "0.2188",
            "yrTerm1.partialUpfront": "0.2233",
            "yrTerm3.partialUpfront": "0.174"
          },
          "ondemand": "0.373"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.156",
            "yrTerm3.allUpfront": "0.082",
            "yrTerm1.allUpfront": "0.1279",
            "yrTerm1.partialUpfront": "0.1305",
            "yrTerm3.partialUpfront": "0.0872"
          },
          "ondemand": "0.233"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.284",
            "yrTerm3.allUpfront": "0.1782",
            "yrTerm1.allUpfront": "0.2378",
            "yrTerm1.partialUpfront": "0.2426",
            "yrTerm3.partialUpfront": "0.1896"
          },
          "ondemand": "0.435"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.665",
            "yrTerm3.allUpfront": "0.4827",
            "yrTerm1.allUpfront": "0.5567",
            "yrTerm1.partialUpfront": "0.5681",
            "yrTerm3.partialUpfront": "0.5135"
          },
          "ondemand": "0.955"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.232",
            "yrTerm3.allUpfront": "0.1467",
            "yrTerm1.allUpfront": "0.1944",
            "yrTerm1.partialUpfront": "0.1983",
            "yrTerm3.partialUpfront": "0.156"
          },
          "ondemand": "0.299"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.112",
            "yrTerm3.allUpfront": "0.0588",
            "yrTerm1.allUpfront": "0.0953",
            "yrTerm1.partialUpfront": "0.0973",
            "yrTerm3.partialUpfront": "0.0625"
          },
          "ondemand": "0.175"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.281",
            "yrTerm3.allUpfront": "0.1777",
            "yrTerm1.allUpfront": "0.2189",
            "yrTerm1.partialUpfront": "0.2406",
            "yrTerm3.partialUpfront": "0.1952"
          },
          "ondemand": "0.517"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.695",
            "yrTerm3.allUpfront": "0.4885",
            "yrTerm1.allUpfront": "0.5406",
            "yrTerm1.partialUpfront": "0.5942",
            "yrTerm3.partialUpfront": "0.5368"
          },
          "ondemand": "1.049"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.256",
            "yrTerm3.allUpfront": "0.1532",
            "yrTerm1.allUpfront": "0.1974",
            "yrTerm1.partialUpfront": "0.219",
            "yrTerm3.partialUpfront": "0.1733"
          },
          "ondemand": "0.335"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.132",
            "yrTerm3.allUpfront": "0.0704",
            "yrTerm1.allUpfront": "0.1043",
            "yrTerm1.partialUpfront": "0.1114",
            "yrTerm3.partialUpfront": "0.077"
          },
          "ondemand": "0.211"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.322",
            "yrTerm3.allUpfront": "0.2025",
            "yrTerm1.allUpfront": "0.2701",
            "yrTerm1.partialUpfront": "0.2756",
            "yrTerm3.partialUpfront": "0.2154"
          },
          "ondemand": "0.47"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.706",
            "yrTerm3.allUpfront": "0.5079",
            "yrTerm1.allUpfront": "0.5916",
            "yrTerm1.partialUpfront": "0.6037",
            "yrTerm3.partialUpfront": "0.5403"
          },
          "ondemand": "0.994"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.265",
            "yrTerm3.allUpfront": "0.1664",
            "yrTerm1.allUpfront": "0.2218",
            "yrTerm1.partialUpfront": "0.2263",
            "yrTerm3.partialUpfront": "0.177"
          },
          "ondemand": "0.314"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.136",
            "yrTerm3.allUpfront": "0.0747",
            "yrTerm1.allUpfront": "0.1178",
            "yrTerm1.partialUpfront": "0.1203",
            "yrTerm3.partialUpfront": "0.0795"
          },
          "ondemand": "0.19"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.284",
            "yrTerm3.allUpfront": "0.1782",
            "yrTerm1.allUpfront": "0.2378",
            "yrTerm1.partialUpfront": "0.2426",
            "yrTerm3.partialUpfront": "0.1896"
          },
          "ondemand": "0.435"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.665",
            "yrTerm3.allUpfront": "0.4827",
            "yrTerm1.allUpfront": "0.5567",
            "yrTerm1.partialUpfront": "0.5681",
            "yrTerm3.partialUpfront": "0.5135"
          },
          "ondemand": "0.955"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.265",
            "yrTerm3.allUpfront": "0.1664",
            "yrTerm1.allUpfront": "0.2218",
            "yrTerm1.partialUpfront": "0.2263",
            "yrTerm3.partialUpfront": "0.177"
          },
          "ondemand": "0.299"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.136",
            "yrTerm3.allUpfront": "0.0747",
            "yrTerm1.allUpfront": "0.1178",
            "yrTerm1.partialUpfront": "0.1203",
            "yrTerm3.partialUpfront": "0.0795"
          },
          "ondemand": "0.19"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 10,
      "max_enis": 3
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "PV"
    ],
    "ebs_optimized": true,
    "storage": {
      "ssd": false,
      "devices": 2,
      "size": 420.0
    },
    "instance_type": "m1.large",
    "ECU": 4.0,
    "memory": 7.5,
    "ebs_max_bandwidth": 500.0
  },
  {
    "family": "General purpose",
    "enhanced_networking": false,
    "vCPU": 4,
    "generation": "previous",
    "ebs_iops": 125.0,
    "network_performance": "High",
    "ebs_throughput": 8000.0,
    "pretty_name": "M1 General Purpose Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.543",
            "yrTerm3.allUpfront": "0.3366",
            "yrTerm1.allUpfront": "0.4553",
            "yrTerm1.partialUpfront": "0.4645",
            "yrTerm3.partialUpfront": "0.3582"
          },
          "ondemand": "0.83"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.906",
            "yrTerm3.allUpfront": "0.6264",
            "yrTerm1.allUpfront": "0.7587",
            "yrTerm1.partialUpfront": "0.7741",
            "yrTerm3.partialUpfront": "0.6664"
          },
          "ondemand": "1.362"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.465",
            "yrTerm3.allUpfront": "0.2909",
            "yrTerm1.allUpfront": "0.389",
            "yrTerm1.partialUpfront": "0.397",
            "yrTerm3.partialUpfront": "0.3095"
          },
          "ondemand": "0.598"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.224",
            "yrTerm3.allUpfront": "0.1167",
            "yrTerm1.allUpfront": "0.1908",
            "yrTerm1.partialUpfront": "0.1947",
            "yrTerm3.partialUpfront": "0.1241"
          },
          "ondemand": "0.35"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.638",
            "yrTerm3.allUpfront": "0.4024",
            "yrTerm1.allUpfront": "0.5341",
            "yrTerm1.partialUpfront": "0.545",
            "yrTerm3.partialUpfront": "0.4281"
          },
          "ondemand": "0.83"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.033",
            "yrTerm3.allUpfront": "0.7201",
            "yrTerm1.allUpfront": "0.8655",
            "yrTerm1.partialUpfront": "0.8832",
            "yrTerm3.partialUpfront": "0.7661"
          },
          "ondemand": "1.435"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.53",
            "yrTerm3.allUpfront": "0.3409",
            "yrTerm1.allUpfront": "0.4437",
            "yrTerm1.partialUpfront": "0.4528",
            "yrTerm3.partialUpfront": "0.3627"
          },
          "ondemand": "0.706"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.296",
            "yrTerm3.allUpfront": "0.1643",
            "yrTerm1.allUpfront": "0.2556",
            "yrTerm1.partialUpfront": "0.2608",
            "yrTerm3.partialUpfront": "0.1748"
          },
          "ondemand": "0.486"
        }
      },
      "sa-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.78",
            "yrTerm3.allUpfront": "0.4448",
            "yrTerm1.allUpfront": "0.6532",
            "yrTerm1.partialUpfront": "0.6665",
            "yrTerm3.partialUpfront": "0.4732"
          },
          "ondemand": "1.067"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.248",
            "yrTerm3.allUpfront": "0.7975",
            "yrTerm1.allUpfront": "1.0455",
            "yrTerm1.partialUpfront": "1.0669",
            "yrTerm3.partialUpfront": "0.8484"
          },
          "ondemand": "1.628"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.661",
            "yrTerm3.allUpfront": "0.3698",
            "yrTerm1.allUpfront": "0.622",
            "yrTerm1.partialUpfront": "0.6347",
            "yrTerm3.partialUpfront": "0.3934"
          },
          "ondemand": "0.715"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.272",
            "yrTerm3.allUpfront": "0.1609",
            "yrTerm1.allUpfront": "0.2273",
            "yrTerm1.partialUpfront": "0.2319",
            "yrTerm3.partialUpfront": "0.1712"
          },
          "ondemand": "0.467"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.622",
            "yrTerm3.allUpfront": "0.3868",
            "yrTerm1.allUpfront": "0.5209",
            "yrTerm1.partialUpfront": "0.5316",
            "yrTerm3.partialUpfront": "0.4115"
          },
          "ondemand": "0.83"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.983",
            "yrTerm3.allUpfront": "0.6766",
            "yrTerm1.allUpfront": "0.8234",
            "yrTerm1.partialUpfront": "0.8402",
            "yrTerm3.partialUpfront": "0.7197"
          },
          "ondemand": "1.362"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.522",
            "yrTerm3.allUpfront": "0.3276",
            "yrTerm1.allUpfront": "0.4371",
            "yrTerm1.partialUpfront": "0.446",
            "yrTerm3.partialUpfront": "0.3485"
          },
          "ondemand": "0.747"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.312",
            "yrTerm3.allUpfront": "0.164",
            "yrTerm1.allUpfront": "0.2568",
            "yrTerm1.partialUpfront": "0.2622",
            "yrTerm3.partialUpfront": "0.1745"
          },
          "ondemand": "0.467"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.622",
            "yrTerm3.allUpfront": "0.3868",
            "yrTerm1.allUpfront": "0.5209",
            "yrTerm1.partialUpfront": "0.5316",
            "yrTerm3.partialUpfront": "0.4115"
          },
          "ondemand": "0.83"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.983",
            "yrTerm3.allUpfront": "0.6766",
            "yrTerm1.allUpfront": "0.8234",
            "yrTerm1.partialUpfront": "0.8402",
            "yrTerm3.partialUpfront": "0.7197"
          },
          "ondemand": "1.362"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.522",
            "yrTerm3.allUpfront": "0.3276",
            "yrTerm1.allUpfront": "0.4371",
            "yrTerm1.partialUpfront": "0.446",
            "yrTerm3.partialUpfront": "0.3485"
          },
          "ondemand": "0.747"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.312",
            "yrTerm3.allUpfront": "0.164",
            "yrTerm1.allUpfront": "0.2568",
            "yrTerm1.partialUpfront": "0.2622",
            "yrTerm3.partialUpfront": "0.1745"
          },
          "ondemand": "0.467"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.543",
            "yrTerm3.allUpfront": "0.3366",
            "yrTerm1.allUpfront": "0.4553",
            "yrTerm1.partialUpfront": "0.4645",
            "yrTerm3.partialUpfront": "0.3582"
          },
          "ondemand": "0.83"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.906",
            "yrTerm3.allUpfront": "0.6264",
            "yrTerm1.allUpfront": "0.7587",
            "yrTerm1.partialUpfront": "0.7741",
            "yrTerm3.partialUpfront": "0.6664"
          },
          "ondemand": "1.362"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.465",
            "yrTerm3.allUpfront": "0.2909",
            "yrTerm1.allUpfront": "0.389",
            "yrTerm1.partialUpfront": "0.397",
            "yrTerm3.partialUpfront": "0.3095"
          },
          "ondemand": "0.598"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.224",
            "yrTerm3.allUpfront": "0.1167",
            "yrTerm1.allUpfront": "0.1908",
            "yrTerm1.partialUpfront": "0.1947",
            "yrTerm3.partialUpfront": "0.1241"
          },
          "ondemand": "0.35"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.679",
            "yrTerm3.allUpfront": "0.3355",
            "yrTerm1.allUpfront": "0.5279",
            "yrTerm1.partialUpfront": "0.5801",
            "yrTerm3.partialUpfront": "0.3687"
          },
          "ondemand": "0.995"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.095",
            "yrTerm3.allUpfront": "0.6487",
            "yrTerm1.allUpfront": "0.8521",
            "yrTerm1.partialUpfront": "0.9362",
            "yrTerm3.partialUpfront": "0.7129"
          },
          "ondemand": "1.545"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.653",
            "yrTerm3.allUpfront": "0.3073",
            "yrTerm1.allUpfront": "0.4978",
            "yrTerm1.partialUpfront": "0.5583",
            "yrTerm3.partialUpfront": "0.3477"
          },
          "ondemand": "0.671"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.264",
            "yrTerm3.allUpfront": "0.1407",
            "yrTerm1.allUpfront": "0.2078",
            "yrTerm1.partialUpfront": "0.2218",
            "yrTerm3.partialUpfront": "0.1539"
          },
          "ondemand": "0.423"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.622",
            "yrTerm3.allUpfront": "0.3868",
            "yrTerm1.allUpfront": "0.5209",
            "yrTerm1.partialUpfront": "0.5316",
            "yrTerm3.partialUpfront": "0.4115"
          },
          "ondemand": "0.899"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.983",
            "yrTerm3.allUpfront": "0.6766",
            "yrTerm1.allUpfront": "0.8234",
            "yrTerm1.partialUpfront": "0.8402",
            "yrTerm3.partialUpfront": "0.7197"
          },
          "ondemand": "1.438"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.529",
            "yrTerm3.allUpfront": "0.3323",
            "yrTerm1.allUpfront": "0.443",
            "yrTerm1.partialUpfront": "0.452",
            "yrTerm3.partialUpfront": "0.3535"
          },
          "ondemand": "0.627"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.272",
            "yrTerm3.allUpfront": "0.1495",
            "yrTerm1.allUpfront": "0.2348",
            "yrTerm1.partialUpfront": "0.2397",
            "yrTerm3.partialUpfront": "0.1591"
          },
          "ondemand": "0.379"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.543",
            "yrTerm3.allUpfront": "0.3366",
            "yrTerm1.allUpfront": "0.4553",
            "yrTerm1.partialUpfront": "0.4645",
            "yrTerm3.partialUpfront": "0.3582"
          },
          "ondemand": "0.83"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.906",
            "yrTerm3.allUpfront": "0.6264",
            "yrTerm1.allUpfront": "0.7587",
            "yrTerm1.partialUpfront": "0.7741",
            "yrTerm3.partialUpfront": "0.6664"
          },
          "ondemand": "1.362"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.529",
            "yrTerm3.allUpfront": "0.3323",
            "yrTerm1.allUpfront": "0.443",
            "yrTerm1.partialUpfront": "0.452",
            "yrTerm3.partialUpfront": "0.3535"
          },
          "ondemand": "0.598"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.272",
            "yrTerm3.allUpfront": "0.1495",
            "yrTerm1.allUpfront": "0.2348",
            "yrTerm1.partialUpfront": "0.2397",
            "yrTerm3.partialUpfront": "0.1591"
          },
          "ondemand": "0.379"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 15,
      "max_enis": 4
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "PV"
    ],
    "ebs_optimized": true,
    "storage": {
      "ssd": false,
      "devices": 4,
      "size": 420.0
    },
    "instance_type": "m1.xlarge",
    "ECU": 8.0,
    "memory": 15.0,
    "ebs_max_bandwidth": 1000.0
  },
  {
    "family": "Compute optimized",
    "enhanced_networking": false,
    "vCPU": 2,
    "generation": "previous",
    "ebs_iops": 0,
    "network_performance": "Moderate",
    "ebs_throughput": 0,
    "pretty_name": "C1 High-CPU Medium",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.215",
            "yrTerm3.allUpfront": "0.1424",
            "yrTerm1.allUpfront": "0.1804",
            "yrTerm1.partialUpfront": "0.1841",
            "yrTerm3.partialUpfront": "0.1515"
          },
          "ondemand": "0.31"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.179",
            "yrTerm3.allUpfront": "0.1157",
            "yrTerm1.allUpfront": "0.1498",
            "yrTerm1.partialUpfront": "0.1529",
            "yrTerm3.partialUpfront": "0.1231"
          },
          "ondemand": "0.21"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.091",
            "yrTerm3.allUpfront": "0.0489",
            "yrTerm1.allUpfront": "0.0766",
            "yrTerm1.partialUpfront": "0.0782",
            "yrTerm3.partialUpfront": "0.052"
          },
          "ondemand": "0.13"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.25",
            "yrTerm3.allUpfront": "0.1693",
            "yrTerm1.allUpfront": "0.2096",
            "yrTerm1.partialUpfront": "0.2138",
            "yrTerm3.partialUpfront": "0.1801"
          },
          "ondemand": "0.31"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.208",
            "yrTerm3.allUpfront": "0.139",
            "yrTerm1.allUpfront": "0.1743",
            "yrTerm1.partialUpfront": "0.1779",
            "yrTerm3.partialUpfront": "0.1479"
          },
          "ondemand": "0.258"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.12",
            "yrTerm3.allUpfront": "0.0714",
            "yrTerm1.allUpfront": "0.1064",
            "yrTerm1.partialUpfront": "0.1086",
            "yrTerm3.partialUpfront": "0.076"
          },
          "ondemand": "0.158"
        }
      },
      "sa-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.266",
            "yrTerm3.allUpfront": "0.1675",
            "yrTerm1.allUpfront": "0.2228",
            "yrTerm1.partialUpfront": "0.2273",
            "yrTerm3.partialUpfront": "0.1782"
          },
          "ondemand": "0.369"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.218",
            "yrTerm3.allUpfront": "0.1252",
            "yrTerm1.allUpfront": "0.1825",
            "yrTerm1.partialUpfront": "0.1863",
            "yrTerm3.partialUpfront": "0.1332"
          },
          "ondemand": "0.259"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.13",
            "yrTerm3.allUpfront": "0.0763",
            "yrTerm1.allUpfront": "0.1096",
            "yrTerm1.partialUpfront": "0.1119",
            "yrTerm3.partialUpfront": "0.0812"
          },
          "ondemand": "0.179"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.242",
            "yrTerm3.allUpfront": "0.1612",
            "yrTerm1.allUpfront": "0.2031",
            "yrTerm1.partialUpfront": "0.2072",
            "yrTerm3.partialUpfront": "0.1715"
          },
          "ondemand": "0.31"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.203",
            "yrTerm3.allUpfront": "0.1336",
            "yrTerm1.allUpfront": "0.1704",
            "yrTerm1.partialUpfront": "0.1739",
            "yrTerm3.partialUpfront": "0.1421"
          },
          "ondemand": "0.266"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.122",
            "yrTerm3.allUpfront": "0.07",
            "yrTerm1.allUpfront": "0.1034",
            "yrTerm1.partialUpfront": "0.1056",
            "yrTerm3.partialUpfront": "0.0745"
          },
          "ondemand": "0.164"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.242",
            "yrTerm3.allUpfront": "0.1612",
            "yrTerm1.allUpfront": "0.2031",
            "yrTerm1.partialUpfront": "0.2072",
            "yrTerm3.partialUpfront": "0.1715"
          },
          "ondemand": "0.31"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.203",
            "yrTerm3.allUpfront": "0.1336",
            "yrTerm1.allUpfront": "0.1704",
            "yrTerm1.partialUpfront": "0.1739",
            "yrTerm3.partialUpfront": "0.1421"
          },
          "ondemand": "0.266"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.122",
            "yrTerm3.allUpfront": "0.07",
            "yrTerm1.allUpfront": "0.1034",
            "yrTerm1.partialUpfront": "0.1056",
            "yrTerm3.partialUpfront": "0.0745"
          },
          "ondemand": "0.164"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.215",
            "yrTerm3.allUpfront": "0.1424",
            "yrTerm1.allUpfront": "0.1804",
            "yrTerm1.partialUpfront": "0.1841",
            "yrTerm3.partialUpfront": "0.1515"
          },
          "ondemand": "0.31"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.179",
            "yrTerm3.allUpfront": "0.1157",
            "yrTerm1.allUpfront": "0.1498",
            "yrTerm1.partialUpfront": "0.1529",
            "yrTerm3.partialUpfront": "0.1231"
          },
          "ondemand": "0.21"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.091",
            "yrTerm3.allUpfront": "0.0489",
            "yrTerm1.allUpfront": "0.0766",
            "yrTerm1.partialUpfront": "0.0782",
            "yrTerm3.partialUpfront": "0.052"
          },
          "ondemand": "0.13"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.227",
            "yrTerm3.allUpfront": "0.1502",
            "yrTerm1.allUpfront": "0.1764",
            "yrTerm1.partialUpfront": "0.1937",
            "yrTerm3.partialUpfront": "0.1651"
          },
          "ondemand": "0.366"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.201",
            "yrTerm3.allUpfront": "0.1259",
            "yrTerm1.allUpfront": "0.1542",
            "yrTerm1.partialUpfront": "0.1719",
            "yrTerm3.partialUpfront": "0.1431"
          },
          "ondemand": "0.237"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.108",
            "yrTerm3.allUpfront": "0.0585",
            "yrTerm1.allUpfront": "0.0868",
            "yrTerm1.partialUpfront": "0.093",
            "yrTerm3.partialUpfront": "0.0642"
          },
          "ondemand": "0.157"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.242",
            "yrTerm3.allUpfront": "0.1612",
            "yrTerm1.allUpfront": "0.2031",
            "yrTerm1.partialUpfront": "0.2072",
            "yrTerm3.partialUpfront": "0.1715"
          },
          "ondemand": "0.329"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.203",
            "yrTerm3.allUpfront": "0.1336",
            "yrTerm1.allUpfront": "0.1704",
            "yrTerm1.partialUpfront": "0.1739",
            "yrTerm3.partialUpfront": "0.1421"
          },
          "ondemand": "0.228"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.111",
            "yrTerm3.allUpfront": "0.0639",
            "yrTerm1.allUpfront": "0.0952",
            "yrTerm1.partialUpfront": "0.0972",
            "yrTerm3.partialUpfront": "0.068"
          },
          "ondemand": "0.148"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.215",
            "yrTerm3.allUpfront": "0.1424",
            "yrTerm1.allUpfront": "0.1804",
            "yrTerm1.partialUpfront": "0.1841",
            "yrTerm3.partialUpfront": "0.1515"
          },
          "ondemand": "0.31"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.203",
            "yrTerm3.allUpfront": "0.1336",
            "yrTerm1.allUpfront": "0.1704",
            "yrTerm1.partialUpfront": "0.1739",
            "yrTerm3.partialUpfront": "0.1421"
          },
          "ondemand": "0.21"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.111",
            "yrTerm3.allUpfront": "0.0639",
            "yrTerm1.allUpfront": "0.0952",
            "yrTerm1.partialUpfront": "0.0972",
            "yrTerm3.partialUpfront": "0.068"
          },
          "ondemand": "0.148"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 6,
      "max_enis": 2
    },
    "arch": [
      "i386",
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "PV"
    ],
    "ebs_optimized": false,
    "storage": {
      "ssd": false,
      "devices": 1,
      "size": 350.0
    },
    "instance_type": "c1.medium",
    "ECU": 5.0,
    "memory": 1.7,
    "ebs_max_bandwidth": 0
  },
  {
    "family": "Compute optimized",
    "enhanced_networking": false,
    "vCPU": 8,
    "generation": "previous",
    "ebs_iops": 125.0,
    "network_performance": "High",
    "ebs_throughput": 8000.0,
    "pretty_name": "C1 High-CPU Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.793",
            "yrTerm3.allUpfront": "0.517",
            "yrTerm1.allUpfront": "0.6643",
            "yrTerm1.partialUpfront": "0.6778",
            "yrTerm3.partialUpfront": "0.55"
          },
          "ondemand": "1.16"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.537",
            "yrTerm3.allUpfront": "1.1142",
            "yrTerm1.allUpfront": "1.2876",
            "yrTerm1.partialUpfront": "1.3138",
            "yrTerm3.partialUpfront": "1.1854"
          },
          "ondemand": "2.124"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.719",
            "yrTerm3.allUpfront": "0.465",
            "yrTerm1.allUpfront": "0.6021",
            "yrTerm1.partialUpfront": "0.6143",
            "yrTerm3.partialUpfront": "0.4947"
          },
          "ondemand": "0.84"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.364",
            "yrTerm3.allUpfront": "0.1955",
            "yrTerm1.allUpfront": "0.3086",
            "yrTerm1.partialUpfront": "0.3149",
            "yrTerm3.partialUpfront": "0.208"
          },
          "ondemand": "0.52"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.926",
            "yrTerm3.allUpfront": "0.619",
            "yrTerm1.allUpfront": "0.7756",
            "yrTerm1.partialUpfront": "0.7914",
            "yrTerm3.partialUpfront": "0.6585"
          },
          "ondemand": "1.16"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.74",
            "yrTerm3.allUpfront": "1.2737",
            "yrTerm1.allUpfront": "1.4578",
            "yrTerm1.partialUpfront": "1.4875",
            "yrTerm3.partialUpfront": "1.355"
          },
          "ondemand": "2.21"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.832",
            "yrTerm3.allUpfront": "0.5551",
            "yrTerm1.allUpfront": "0.6966",
            "yrTerm1.partialUpfront": "0.7108",
            "yrTerm3.partialUpfront": "0.5905"
          },
          "ondemand": "1.032"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.48",
            "yrTerm3.allUpfront": "0.2874",
            "yrTerm1.allUpfront": "0.4265",
            "yrTerm1.partialUpfront": "0.4353",
            "yrTerm3.partialUpfront": "0.3057"
          },
          "ondemand": "0.632"
        }
      },
      "sa-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.942",
            "yrTerm3.allUpfront": "0.5861",
            "yrTerm1.allUpfront": "0.7888",
            "yrTerm1.partialUpfront": "0.8049",
            "yrTerm3.partialUpfront": "0.6235"
          },
          "ondemand": "1.398"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.77",
            "yrTerm3.allUpfront": "1.2393",
            "yrTerm1.allUpfront": "1.4828",
            "yrTerm1.partialUpfront": "1.513",
            "yrTerm3.partialUpfront": "1.3184"
          },
          "ondemand": "2.378"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.871",
            "yrTerm3.allUpfront": "0.4602",
            "yrTerm1.allUpfront": "0.7295",
            "yrTerm1.partialUpfront": "0.7443",
            "yrTerm3.partialUpfront": "0.4896"
          },
          "ondemand": "1.038"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.52",
            "yrTerm3.allUpfront": "0.3051",
            "yrTerm1.allUpfront": "0.4366",
            "yrTerm1.partialUpfront": "0.4456",
            "yrTerm3.partialUpfront": "0.3246"
          },
          "ondemand": "0.718"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.897",
            "yrTerm3.allUpfront": "0.5906",
            "yrTerm1.allUpfront": "0.7511",
            "yrTerm1.partialUpfront": "0.7665",
            "yrTerm3.partialUpfront": "0.6283"
          },
          "ondemand": "1.16"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.641",
            "yrTerm3.allUpfront": "1.1888",
            "yrTerm1.allUpfront": "1.3744",
            "yrTerm1.partialUpfront": "1.4025",
            "yrTerm3.partialUpfront": "1.2647"
          },
          "ondemand": "2.124"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.812",
            "yrTerm3.allUpfront": "0.5327",
            "yrTerm1.allUpfront": "0.6805",
            "yrTerm1.partialUpfront": "0.6943",
            "yrTerm3.partialUpfront": "0.5667"
          },
          "ondemand": "1.065"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.488",
            "yrTerm3.allUpfront": "0.278",
            "yrTerm1.allUpfront": "0.4137",
            "yrTerm1.partialUpfront": "0.4222",
            "yrTerm3.partialUpfront": "0.2958"
          },
          "ondemand": "0.655"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.897",
            "yrTerm3.allUpfront": "0.5906",
            "yrTerm1.allUpfront": "0.7511",
            "yrTerm1.partialUpfront": "0.7665",
            "yrTerm3.partialUpfront": "0.6283"
          },
          "ondemand": "1.16"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.641",
            "yrTerm3.allUpfront": "1.1888",
            "yrTerm1.allUpfront": "1.3744",
            "yrTerm1.partialUpfront": "1.4025",
            "yrTerm3.partialUpfront": "1.2647"
          },
          "ondemand": "2.124"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.812",
            "yrTerm3.allUpfront": "0.5327",
            "yrTerm1.allUpfront": "0.6805",
            "yrTerm1.partialUpfront": "0.6943",
            "yrTerm3.partialUpfront": "0.5667"
          },
          "ondemand": "1.065"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.488",
            "yrTerm3.allUpfront": "0.278",
            "yrTerm1.allUpfront": "0.4137",
            "yrTerm1.partialUpfront": "0.4222",
            "yrTerm3.partialUpfront": "0.2958"
          },
          "ondemand": "0.655"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.793",
            "yrTerm3.allUpfront": "0.517",
            "yrTerm1.allUpfront": "0.6643",
            "yrTerm1.partialUpfront": "0.6778",
            "yrTerm3.partialUpfront": "0.55"
          },
          "ondemand": "1.16"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.537",
            "yrTerm3.allUpfront": "1.1142",
            "yrTerm1.allUpfront": "1.2876",
            "yrTerm1.partialUpfront": "1.3138",
            "yrTerm3.partialUpfront": "1.1854"
          },
          "ondemand": "2.124"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.719",
            "yrTerm3.allUpfront": "0.465",
            "yrTerm1.allUpfront": "0.6021",
            "yrTerm1.partialUpfront": "0.6143",
            "yrTerm3.partialUpfront": "0.4947"
          },
          "ondemand": "0.84"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.364",
            "yrTerm3.allUpfront": "0.1955",
            "yrTerm1.allUpfront": "0.3086",
            "yrTerm1.partialUpfront": "0.3149",
            "yrTerm3.partialUpfront": "0.208"
          },
          "ondemand": "0.52"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.86",
            "yrTerm3.allUpfront": "0.5613",
            "yrTerm1.allUpfront": "0.6686",
            "yrTerm1.partialUpfront": "0.7347",
            "yrTerm3.partialUpfront": "0.6168"
          },
          "ondemand": "1.384"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.635",
            "yrTerm3.allUpfront": "1.154",
            "yrTerm1.allUpfront": "1.2713",
            "yrTerm1.partialUpfront": "1.3972",
            "yrTerm3.partialUpfront": "1.2682"
          },
          "ondemand": "2.331"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.809",
            "yrTerm3.allUpfront": "0.5043",
            "yrTerm1.allUpfront": "0.6199",
            "yrTerm1.partialUpfront": "0.6911",
            "yrTerm3.partialUpfront": "0.5735"
          },
          "ondemand": "0.948"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.432",
            "yrTerm3.allUpfront": "0.2331",
            "yrTerm1.allUpfront": "0.3489",
            "yrTerm1.partialUpfront": "0.3743",
            "yrTerm3.partialUpfront": "0.2558"
          },
          "ondemand": "0.628"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.897",
            "yrTerm3.allUpfront": "0.5906",
            "yrTerm1.allUpfront": "0.7511",
            "yrTerm1.partialUpfront": "0.7665",
            "yrTerm3.partialUpfront": "0.6283"
          },
          "ondemand": "1.236"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.641",
            "yrTerm3.allUpfront": "1.1888",
            "yrTerm1.allUpfront": "1.3744",
            "yrTerm1.partialUpfront": "1.4025",
            "yrTerm3.partialUpfront": "1.2647"
          },
          "ondemand": "2.206"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.815",
            "yrTerm3.allUpfront": "0.5336",
            "yrTerm1.allUpfront": "0.6824",
            "yrTerm1.partialUpfront": "0.6963",
            "yrTerm3.partialUpfront": "0.5677"
          },
          "ondemand": "0.912"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.444",
            "yrTerm3.allUpfront": "0.2557",
            "yrTerm1.allUpfront": "0.3821",
            "yrTerm1.partialUpfront": "0.3899",
            "yrTerm3.partialUpfront": "0.272"
          },
          "ondemand": "0.592"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.793",
            "yrTerm3.allUpfront": "0.517",
            "yrTerm1.allUpfront": "0.6643",
            "yrTerm1.partialUpfront": "0.6778",
            "yrTerm3.partialUpfront": "0.55"
          },
          "ondemand": "1.16"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.537",
            "yrTerm3.allUpfront": "1.1142",
            "yrTerm1.allUpfront": "1.2876",
            "yrTerm1.partialUpfront": "1.3138",
            "yrTerm3.partialUpfront": "1.1854"
          },
          "ondemand": "2.124"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.815",
            "yrTerm3.allUpfront": "0.5336",
            "yrTerm1.allUpfront": "0.6824",
            "yrTerm1.partialUpfront": "0.6963",
            "yrTerm3.partialUpfront": "0.5677"
          },
          "ondemand": "0.84"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.444",
            "yrTerm3.allUpfront": "0.2557",
            "yrTerm1.allUpfront": "0.3821",
            "yrTerm1.partialUpfront": "0.3899",
            "yrTerm3.partialUpfront": "0.272"
          },
          "ondemand": "0.592"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 15,
      "max_enis": 4
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "PV"
    ],
    "ebs_optimized": true,
    "storage": {
      "ssd": false,
      "devices": 4,
      "size": 420.0
    },
    "instance_type": "c1.xlarge",
    "ECU": 20.0,
    "memory": 7.0,
    "ebs_max_bandwidth": 1000.0
  },
  {
    "family": "Compute optimized",
    "enhanced_networking": false,
    "vCPU": 32,
    "generation": "previous",
    "ebs_iops": 0,
    "network_performance": "10 Gigabit",
    "ebs_throughput": 0,
    "pretty_name": "Cluster Compute Eight Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.437",
            "yrTerm3.allUpfront": "0.8925",
            "yrTerm1.allUpfront": "1.2035",
            "yrTerm1.partialUpfront": "1.2282",
            "yrTerm3.partialUpfront": "0.9495"
          },
          "ondemand": "2.775"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.956",
            "yrTerm3.allUpfront": "2.1119",
            "yrTerm1.allUpfront": "2.476",
            "yrTerm1.partialUpfront": "2.5266",
            "yrTerm3.partialUpfront": "2.2467"
          },
          "ondemand": "4.41"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.336",
            "yrTerm3.allUpfront": "0.8111",
            "yrTerm1.allUpfront": "1.1189",
            "yrTerm1.partialUpfront": "1.1418",
            "yrTerm3.partialUpfront": "0.8629"
          },
          "ondemand": "2.57"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.09",
            "yrTerm3.allUpfront": "0.6137",
            "yrTerm1.allUpfront": "0.9131",
            "yrTerm1.partialUpfront": "0.9318",
            "yrTerm3.partialUpfront": "0.6529"
          },
          "ondemand": "2"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "2.376",
            "yrTerm3.allUpfront": "1.5949",
            "yrTerm1.allUpfront": "1.9905",
            "yrTerm1.partialUpfront": "2.0312",
            "yrTerm3.partialUpfront": "1.6967"
          },
          "ondemand": "3.124"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.753",
            "yrTerm3.allUpfront": "2.7",
            "yrTerm1.allUpfront": "3.1434",
            "yrTerm1.partialUpfront": "3.2075",
            "yrTerm3.partialUpfront": "2.8723"
          },
          "ondemand": "4.733"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.148",
            "yrTerm3.allUpfront": "1.4109",
            "yrTerm1.allUpfront": "1.799",
            "yrTerm1.partialUpfront": "1.8357",
            "yrTerm3.partialUpfront": "1.501"
          },
          "ondemand": "2.919"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.512",
            "yrTerm3.allUpfront": "0.9005",
            "yrTerm1.allUpfront": "1.2668",
            "yrTerm1.partialUpfront": "1.2927",
            "yrTerm3.partialUpfront": "0.958"
          },
          "ondemand": "2.349"
        }
      },
      "sa-east-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.437",
            "yrTerm3.allUpfront": "0.8925",
            "yrTerm1.allUpfront": "1.2035",
            "yrTerm1.partialUpfront": "1.2282",
            "yrTerm3.partialUpfront": "0.9495"
          },
          "ondemand": "2.775"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.956",
            "yrTerm3.allUpfront": "2.1119",
            "yrTerm1.allUpfront": "2.476",
            "yrTerm1.partialUpfront": "2.5266",
            "yrTerm3.partialUpfront": "2.2467"
          },
          "ondemand": "4.41"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.336",
            "yrTerm3.allUpfront": "0.8111",
            "yrTerm1.allUpfront": "1.1189",
            "yrTerm1.partialUpfront": "1.1418",
            "yrTerm3.partialUpfront": "0.8629"
          },
          "ondemand": "2.57"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.09",
            "yrTerm3.allUpfront": "0.6137",
            "yrTerm1.allUpfront": "0.9131",
            "yrTerm1.partialUpfront": "0.9318",
            "yrTerm3.partialUpfront": "0.6529"
          },
          "ondemand": "2"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "2.665",
            "yrTerm3.allUpfront": "1.0797",
            "yrTerm1.allUpfront": "1.397",
            "yrTerm1.partialUpfront": "1.5352",
            "yrTerm3.partialUpfront": "2.3725"
          },
          "ondemand": "2.775"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.694",
            "yrTerm3.allUpfront": "2.5559",
            "yrTerm1.allUpfront": "2.8733",
            "yrTerm1.partialUpfront": "3.1576",
            "yrTerm3.partialUpfront": "2.8087"
          },
          "ondemand": "4.9"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.668",
            "yrTerm3.allUpfront": "0.9591",
            "yrTerm1.allUpfront": "1.2975",
            "yrTerm1.partialUpfront": "1.4258",
            "yrTerm3.partialUpfront": "1.0769"
          },
          "ondemand": "2.57"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.381",
            "yrTerm3.allUpfront": "0.8027",
            "yrTerm1.allUpfront": "1.0893",
            "yrTerm1.partialUpfront": "1.1808",
            "yrTerm3.partialUpfront": "0.9019"
          },
          "ondemand": "2.25"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "2.711",
            "yrTerm3.allUpfront": "1.9487",
            "yrTerm1.allUpfront": "2.6249",
            "yrTerm1.partialUpfront": "2.6784",
            "yrTerm3.partialUpfront": "2.0731"
          },
          "ondemand": "2.775"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.694",
            "yrTerm3.allUpfront": "2.6401",
            "yrTerm1.allUpfront": "3.0944",
            "yrTerm1.partialUpfront": "3.1576",
            "yrTerm3.partialUpfront": "2.8087"
          },
          "ondemand": "4.85"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.668",
            "yrTerm3.allUpfront": "1.0123",
            "yrTerm1.allUpfront": "1.3973",
            "yrTerm1.partialUpfront": "1.4258",
            "yrTerm3.partialUpfront": "1.0769"
          },
          "ondemand": "2.57"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.381",
            "yrTerm3.allUpfront": "0.8477",
            "yrTerm1.allUpfront": "1.1571",
            "yrTerm1.partialUpfront": "1.1808",
            "yrTerm3.partialUpfront": "0.9019"
          },
          "ondemand": "2.25"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 30,
      "max_enis": 8
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": false,
    "storage": {
      "ssd": false,
      "devices": 4,
      "size": 840.0
    },
    "instance_type": "cc2.8xlarge",
    "ECU": 88.0,
    "memory": 60.5,
    "ebs_max_bandwidth": 0
  },
  {
    "family": "GPU instances",
    "enhanced_networking": false,
    "vCPU": 16,
    "generation": "previous",
    "ebs_iops": 0,
    "network_performance": "10 Gigabit",
    "ebs_throughput": 0,
    "pretty_name": "Cluster GPU Quadruple Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "ondemand": "2.703"
        },
        "mswinSQL": {
          "ondemand": "3.85"
        },
        "mswin": {
          "ondemand": "2.6"
        },
        "linux": {
          "ondemand": "2.1"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "sa-east-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "ondemand": "2.703"
        },
        "mswinSQL": {
          "ondemand": "3.85"
        },
        "mswin": {
          "ondemand": "2.6"
        },
        "linux": {
          "ondemand": "2.36"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 30,
      "max_enis": 8
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": false,
    "storage": {
      "ssd": false,
      "devices": 2,
      "size": 840.0
    },
    "instance_type": "cg1.4xlarge",
    "ECU": 33.5,
    "memory": 22.5,
    "ebs_max_bandwidth": 0
  },
  {
    "family": "Memory optimized",
    "enhanced_networking": false,
    "vCPU": 2,
    "generation": "previous",
    "ebs_iops": 0,
    "network_performance": "Moderate",
    "ebs_throughput": 0,
    "pretty_name": "M2 High Memory Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.323",
            "yrTerm3.allUpfront": "0.1928",
            "yrTerm1.allUpfront": "0.2709",
            "yrTerm1.partialUpfront": "0.2764",
            "yrTerm3.partialUpfront": "0.2051"
          },
          "ondemand": "0.445"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.728",
            "yrTerm3.allUpfront": "0.5123",
            "yrTerm1.allUpfront": "0.6095",
            "yrTerm1.partialUpfront": "0.6219",
            "yrTerm3.partialUpfront": "0.545"
          },
          "ondemand": "1.084"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.249",
            "yrTerm3.allUpfront": "0.1495",
            "yrTerm1.allUpfront": "0.2084",
            "yrTerm1.partialUpfront": "0.2128",
            "yrTerm3.partialUpfront": "0.1591"
          },
          "ondemand": "0.345"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.111",
            "yrTerm3.allUpfront": "0.0558",
            "yrTerm1.allUpfront": "0.093",
            "yrTerm1.partialUpfront": "0.095",
            "yrTerm3.partialUpfront": "0.0594"
          },
          "ondemand": "0.245"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.41",
            "yrTerm3.allUpfront": "0.2508",
            "yrTerm1.allUpfront": "0.3435",
            "yrTerm1.partialUpfront": "0.3505",
            "yrTerm3.partialUpfront": "0.2668"
          },
          "ondemand": "0.445"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.857",
            "yrTerm3.allUpfront": "0.6043",
            "yrTerm1.allUpfront": "0.7178",
            "yrTerm1.partialUpfront": "0.7325",
            "yrTerm3.partialUpfront": "0.6429"
          },
          "ondemand": "1.106"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.312",
            "yrTerm3.allUpfront": "0.1936",
            "yrTerm1.allUpfront": "0.2612",
            "yrTerm1.partialUpfront": "0.2665",
            "yrTerm3.partialUpfront": "0.206"
          },
          "ondemand": "0.352"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.151",
            "yrTerm3.allUpfront": "0.0815",
            "yrTerm1.allUpfront": "0.1267",
            "yrTerm1.partialUpfront": "0.1294",
            "yrTerm3.partialUpfront": "0.0868"
          },
          "ondemand": "0.287"
        }
      },
      "sa-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.507",
            "yrTerm3.allUpfront": "0.2811",
            "yrTerm1.allUpfront": "0.425",
            "yrTerm1.partialUpfront": "0.4336",
            "yrTerm3.partialUpfront": "0.299"
          },
          "ondemand": "0.623"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.954",
            "yrTerm3.allUpfront": "0.6234",
            "yrTerm1.allUpfront": "0.7992",
            "yrTerm1.partialUpfront": "0.8155",
            "yrTerm3.partialUpfront": "0.6632"
          },
          "ondemand": "1.314"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.405",
            "yrTerm3.allUpfront": "0.2326",
            "yrTerm1.allUpfront": "0.3882",
            "yrTerm1.partialUpfront": "0.3962",
            "yrTerm3.partialUpfront": "0.2475"
          },
          "ondemand": "0.423"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.129",
            "yrTerm3.allUpfront": "0.0771",
            "yrTerm1.allUpfront": "0.1083",
            "yrTerm1.partialUpfront": "0.1106",
            "yrTerm3.partialUpfront": "0.082"
          },
          "ondemand": "0.323"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.4",
            "yrTerm3.allUpfront": "0.2432",
            "yrTerm1.allUpfront": "0.3352",
            "yrTerm1.partialUpfront": "0.342",
            "yrTerm3.partialUpfront": "0.2587"
          },
          "ondemand": "0.445"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.809",
            "yrTerm3.allUpfront": "0.567",
            "yrTerm1.allUpfront": "0.6774",
            "yrTerm1.partialUpfront": "0.6913",
            "yrTerm3.partialUpfront": "0.6032"
          },
          "ondemand": "1.084"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.306",
            "yrTerm3.allUpfront": "0.189",
            "yrTerm1.allUpfront": "0.2565",
            "yrTerm1.partialUpfront": "0.2618",
            "yrTerm3.partialUpfront": "0.2011"
          },
          "ondemand": "0.371"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.148",
            "yrTerm3.allUpfront": "0.08",
            "yrTerm1.allUpfront": "0.1247",
            "yrTerm1.partialUpfront": "0.1273",
            "yrTerm3.partialUpfront": "0.0851"
          },
          "ondemand": "0.296"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.4",
            "yrTerm3.allUpfront": "0.2432",
            "yrTerm1.allUpfront": "0.3352",
            "yrTerm1.partialUpfront": "0.342",
            "yrTerm3.partialUpfront": "0.2587"
          },
          "ondemand": "0.445"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.809",
            "yrTerm3.allUpfront": "0.567",
            "yrTerm1.allUpfront": "0.6774",
            "yrTerm1.partialUpfront": "0.6913",
            "yrTerm3.partialUpfront": "0.6032"
          },
          "ondemand": "1.084"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.306",
            "yrTerm3.allUpfront": "0.189",
            "yrTerm1.allUpfront": "0.2565",
            "yrTerm1.partialUpfront": "0.2618",
            "yrTerm3.partialUpfront": "0.2011"
          },
          "ondemand": "0.371"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.148",
            "yrTerm3.allUpfront": "0.08",
            "yrTerm1.allUpfront": "0.1247",
            "yrTerm1.partialUpfront": "0.1273",
            "yrTerm3.partialUpfront": "0.0851"
          },
          "ondemand": "0.296"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.323",
            "yrTerm3.allUpfront": "0.1928",
            "yrTerm1.allUpfront": "0.2709",
            "yrTerm1.partialUpfront": "0.2764",
            "yrTerm3.partialUpfront": "0.2051"
          },
          "ondemand": "0.445"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.728",
            "yrTerm3.allUpfront": "0.5123",
            "yrTerm1.allUpfront": "0.6095",
            "yrTerm1.partialUpfront": "0.6219",
            "yrTerm3.partialUpfront": "0.545"
          },
          "ondemand": "1.084"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.249",
            "yrTerm3.allUpfront": "0.1495",
            "yrTerm1.allUpfront": "0.2084",
            "yrTerm1.partialUpfront": "0.2128",
            "yrTerm3.partialUpfront": "0.1591"
          },
          "ondemand": "0.345"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.111",
            "yrTerm3.allUpfront": "0.0558",
            "yrTerm1.allUpfront": "0.093",
            "yrTerm1.partialUpfront": "0.095",
            "yrTerm3.partialUpfront": "0.0594"
          },
          "ondemand": "0.245"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.295",
            "yrTerm3.allUpfront": "0.1817",
            "yrTerm1.allUpfront": "0.2295",
            "yrTerm1.partialUpfront": "0.2521",
            "yrTerm3.partialUpfront": "0.1997"
          },
          "ondemand": "0.533"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.766",
            "yrTerm3.allUpfront": "0.5215",
            "yrTerm1.allUpfront": "0.5954",
            "yrTerm1.partialUpfront": "0.6544",
            "yrTerm3.partialUpfront": "0.5731"
          },
          "ondemand": "1.204"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.27",
            "yrTerm3.allUpfront": "0.1584",
            "yrTerm1.allUpfront": "0.2095",
            "yrTerm1.partialUpfront": "0.2309",
            "yrTerm3.partialUpfront": "0.1782"
          },
          "ondemand": "0.393"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.136",
            "yrTerm3.allUpfront": "0.0679",
            "yrTerm1.allUpfront": "0.1084",
            "yrTerm1.partialUpfront": "0.1168",
            "yrTerm3.partialUpfront": "0.0745"
          },
          "ondemand": "0.293"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.4",
            "yrTerm3.allUpfront": "0.2432",
            "yrTerm1.allUpfront": "0.3352",
            "yrTerm1.partialUpfront": "0.342",
            "yrTerm3.partialUpfront": "0.2587"
          },
          "ondemand": "0.481"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.809",
            "yrTerm3.allUpfront": "0.567",
            "yrTerm1.allUpfront": "0.6774",
            "yrTerm1.partialUpfront": "0.6913",
            "yrTerm3.partialUpfront": "0.6032"
          },
          "ondemand": "1.14"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.313",
            "yrTerm3.allUpfront": "0.1928",
            "yrTerm1.allUpfront": "0.2624",
            "yrTerm1.partialUpfront": "0.2678",
            "yrTerm3.partialUpfront": "0.2051"
          },
          "ondemand": "0.375"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.134",
            "yrTerm3.allUpfront": "0.0718",
            "yrTerm1.allUpfront": "0.1127",
            "yrTerm1.partialUpfront": "0.115",
            "yrTerm3.partialUpfront": "0.0764"
          },
          "ondemand": "0.275"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.323",
            "yrTerm3.allUpfront": "0.1928",
            "yrTerm1.allUpfront": "0.2709",
            "yrTerm1.partialUpfront": "0.2764",
            "yrTerm3.partialUpfront": "0.2051"
          },
          "ondemand": "0.445"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.728",
            "yrTerm3.allUpfront": "0.5123",
            "yrTerm1.allUpfront": "0.6095",
            "yrTerm1.partialUpfront": "0.6219",
            "yrTerm3.partialUpfront": "0.545"
          },
          "ondemand": "1.084"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.313",
            "yrTerm3.allUpfront": "0.1928",
            "yrTerm1.allUpfront": "0.2624",
            "yrTerm1.partialUpfront": "0.2678",
            "yrTerm3.partialUpfront": "0.2051"
          },
          "ondemand": "0.345"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.134",
            "yrTerm3.allUpfront": "0.0727",
            "yrTerm1.allUpfront": "0.1127",
            "yrTerm1.partialUpfront": "0.115",
            "yrTerm3.partialUpfront": "0.0774"
          },
          "ondemand": "0.275"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 15,
      "max_enis": 4
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "PV"
    ],
    "ebs_optimized": false,
    "storage": {
      "ssd": false,
      "devices": 1,
      "size": 420.0
    },
    "instance_type": "m2.xlarge",
    "ECU": 6.5,
    "memory": 17.1,
    "ebs_max_bandwidth": 0
  },
  {
    "family": "Memory optimized",
    "enhanced_networking": false,
    "vCPU": 4,
    "generation": "previous",
    "ebs_iops": 62.5,
    "network_performance": "Moderate",
    "ebs_throughput": 4000.0,
    "pretty_name": "M2 High Memory Double Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.622",
            "yrTerm3.allUpfront": "0.3666",
            "yrTerm1.allUpfront": "0.5207",
            "yrTerm1.partialUpfront": "0.5313",
            "yrTerm3.partialUpfront": "0.3901"
          },
          "ondemand": "0.85"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.944",
            "yrTerm3.allUpfront": "0.6287",
            "yrTerm1.allUpfront": "0.7905",
            "yrTerm1.partialUpfront": "0.8066",
            "yrTerm3.partialUpfront": "0.6688"
          },
          "ondemand": "1.489"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.496",
            "yrTerm3.allUpfront": "0.3006",
            "yrTerm1.allUpfront": "0.4158",
            "yrTerm1.partialUpfront": "0.4243",
            "yrTerm3.partialUpfront": "0.3198"
          },
          "ondemand": "0.69"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.222",
            "yrTerm3.allUpfront": "0.1116",
            "yrTerm1.allUpfront": "0.1863",
            "yrTerm1.partialUpfront": "0.1901",
            "yrTerm3.partialUpfront": "0.1187"
          },
          "ondemand": "0.49"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.794",
            "yrTerm3.allUpfront": "0.4821",
            "yrTerm1.allUpfront": "0.6653",
            "yrTerm1.partialUpfront": "0.6789",
            "yrTerm3.partialUpfront": "0.5129"
          },
          "ondemand": "0.85"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.144",
            "yrTerm3.allUpfront": "0.7675",
            "yrTerm1.allUpfront": "0.9579",
            "yrTerm1.partialUpfront": "0.9774",
            "yrTerm3.partialUpfront": "0.8165"
          },
          "ondemand": "1.558"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.624",
            "yrTerm3.allUpfront": "0.3875",
            "yrTerm1.allUpfront": "0.5226",
            "yrTerm1.partialUpfront": "0.5333",
            "yrTerm3.partialUpfront": "0.4122"
          },
          "ondemand": "0.705"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.302",
            "yrTerm3.allUpfront": "0.1632",
            "yrTerm1.allUpfront": "0.2547",
            "yrTerm1.partialUpfront": "0.2599",
            "yrTerm3.partialUpfront": "0.1736"
          },
          "ondemand": "0.575"
        }
      },
      "sa-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.992",
            "yrTerm3.allUpfront": "0.5441",
            "yrTerm1.allUpfront": "0.8308",
            "yrTerm1.partialUpfront": "0.8477",
            "yrTerm3.partialUpfront": "0.5788"
          },
          "ondemand": "1.205"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.353",
            "yrTerm3.allUpfront": "0.8288",
            "yrTerm1.allUpfront": "1.1332",
            "yrTerm1.partialUpfront": "1.1563",
            "yrTerm3.partialUpfront": "0.8817"
          },
          "ondemand": "1.903"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.809",
            "yrTerm3.allUpfront": "0.4648",
            "yrTerm1.allUpfront": "0.7747",
            "yrTerm1.partialUpfront": "0.7905",
            "yrTerm3.partialUpfront": "0.4945"
          },
          "ondemand": "0.845"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.258",
            "yrTerm3.allUpfront": "0.1551",
            "yrTerm1.allUpfront": "0.2176",
            "yrTerm1.partialUpfront": "0.2221",
            "yrTerm3.partialUpfront": "0.165"
          },
          "ondemand": "0.645"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.775",
            "yrTerm3.allUpfront": "0.4683",
            "yrTerm1.allUpfront": "0.6491",
            "yrTerm1.partialUpfront": "0.6624",
            "yrTerm3.partialUpfront": "0.4982"
          },
          "ondemand": "0.85"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.092",
            "yrTerm3.allUpfront": "0.7272",
            "yrTerm1.allUpfront": "0.9143",
            "yrTerm1.partialUpfront": "0.9329",
            "yrTerm3.partialUpfront": "0.7736"
          },
          "ondemand": "1.654"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.613",
            "yrTerm3.allUpfront": "0.3787",
            "yrTerm1.allUpfront": "0.5138",
            "yrTerm1.partialUpfront": "0.5243",
            "yrTerm3.partialUpfront": "0.4028"
          },
          "ondemand": "0.742"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.296",
            "yrTerm3.allUpfront": "0.16",
            "yrTerm1.allUpfront": "0.2502",
            "yrTerm1.partialUpfront": "0.2554",
            "yrTerm3.partialUpfront": "0.1702"
          },
          "ondemand": "0.592"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.775",
            "yrTerm3.allUpfront": "0.4683",
            "yrTerm1.allUpfront": "0.6491",
            "yrTerm1.partialUpfront": "0.6624",
            "yrTerm3.partialUpfront": "0.4982"
          },
          "ondemand": "0.85"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.092",
            "yrTerm3.allUpfront": "0.7272",
            "yrTerm1.allUpfront": "0.9143",
            "yrTerm1.partialUpfront": "0.9329",
            "yrTerm3.partialUpfront": "0.7736"
          },
          "ondemand": "1.654"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.613",
            "yrTerm3.allUpfront": "0.3787",
            "yrTerm1.allUpfront": "0.5138",
            "yrTerm1.partialUpfront": "0.5243",
            "yrTerm3.partialUpfront": "0.4028"
          },
          "ondemand": "0.742"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.296",
            "yrTerm3.allUpfront": "0.16",
            "yrTerm1.allUpfront": "0.2502",
            "yrTerm1.partialUpfront": "0.2554",
            "yrTerm3.partialUpfront": "0.1702"
          },
          "ondemand": "0.592"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.622",
            "yrTerm3.allUpfront": "0.3666",
            "yrTerm1.allUpfront": "0.5207",
            "yrTerm1.partialUpfront": "0.5313",
            "yrTerm3.partialUpfront": "0.3901"
          },
          "ondemand": "0.85"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.944",
            "yrTerm3.allUpfront": "0.6287",
            "yrTerm1.allUpfront": "0.7905",
            "yrTerm1.partialUpfront": "0.8066",
            "yrTerm3.partialUpfront": "0.6688"
          },
          "ondemand": "1.489"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.496",
            "yrTerm3.allUpfront": "0.3006",
            "yrTerm1.allUpfront": "0.4158",
            "yrTerm1.partialUpfront": "0.4243",
            "yrTerm3.partialUpfront": "0.3198"
          },
          "ondemand": "0.69"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.222",
            "yrTerm3.allUpfront": "0.1116",
            "yrTerm1.allUpfront": "0.1863",
            "yrTerm1.partialUpfront": "0.1901",
            "yrTerm3.partialUpfront": "0.1187"
          },
          "ondemand": "0.49"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.566",
            "yrTerm3.allUpfront": "0.3427",
            "yrTerm1.allUpfront": "0.4398",
            "yrTerm1.partialUpfront": "0.4834",
            "yrTerm3.partialUpfront": "0.3766"
          },
          "ondemand": "1.026"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.015",
            "yrTerm3.allUpfront": "0.6552",
            "yrTerm1.allUpfront": "0.7893",
            "yrTerm1.partialUpfront": "0.8673",
            "yrTerm3.partialUpfront": "0.72"
          },
          "ondemand": "1.705"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.54",
            "yrTerm3.allUpfront": "0.3158",
            "yrTerm1.allUpfront": "0.4189",
            "yrTerm1.partialUpfront": "0.4618",
            "yrTerm3.partialUpfront": "0.3552"
          },
          "ondemand": "0.786"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.272",
            "yrTerm3.allUpfront": "0.1358",
            "yrTerm1.allUpfront": "0.2177",
            "yrTerm1.partialUpfront": "0.2347",
            "yrTerm3.partialUpfront": "0.149"
          },
          "ondemand": "0.586"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.775",
            "yrTerm3.allUpfront": "0.4683",
            "yrTerm1.allUpfront": "0.6491",
            "yrTerm1.partialUpfront": "0.6624",
            "yrTerm3.partialUpfront": "0.4982"
          },
          "ondemand": "0.922"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.092",
            "yrTerm3.allUpfront": "0.7272",
            "yrTerm1.allUpfront": "0.9143",
            "yrTerm1.partialUpfront": "0.9329",
            "yrTerm3.partialUpfront": "0.7736"
          },
          "ondemand": "1.589"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.625",
            "yrTerm3.allUpfront": "0.3871",
            "yrTerm1.allUpfront": "0.5236",
            "yrTerm1.partialUpfront": "0.5343",
            "yrTerm3.partialUpfront": "0.4118"
          },
          "ondemand": "0.75"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.268",
            "yrTerm3.allUpfront": "0.1435",
            "yrTerm1.allUpfront": "0.2255",
            "yrTerm1.partialUpfront": "0.2301",
            "yrTerm3.partialUpfront": "0.1527"
          },
          "ondemand": "0.55"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.622",
            "yrTerm3.allUpfront": "0.3666",
            "yrTerm1.allUpfront": "0.5207",
            "yrTerm1.partialUpfront": "0.5313",
            "yrTerm3.partialUpfront": "0.3901"
          },
          "ondemand": "0.85"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.944",
            "yrTerm3.allUpfront": "0.6287",
            "yrTerm1.allUpfront": "0.7905",
            "yrTerm1.partialUpfront": "0.8066",
            "yrTerm3.partialUpfront": "0.6688"
          },
          "ondemand": "1.489"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.625",
            "yrTerm3.allUpfront": "0.3871",
            "yrTerm1.allUpfront": "0.5236",
            "yrTerm1.partialUpfront": "0.5343",
            "yrTerm3.partialUpfront": "0.4118"
          },
          "ondemand": "0.69"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.268",
            "yrTerm3.allUpfront": "0.1454",
            "yrTerm1.allUpfront": "0.2255",
            "yrTerm1.partialUpfront": "0.2301",
            "yrTerm3.partialUpfront": "0.1547"
          },
          "ondemand": "0.55"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 30,
      "max_enis": 4
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "PV"
    ],
    "ebs_optimized": true,
    "storage": {
      "ssd": false,
      "devices": 1,
      "size": 850.0
    },
    "instance_type": "m2.2xlarge",
    "ECU": 13.0,
    "memory": 34.2,
    "ebs_max_bandwidth": 500.0
  },
  {
    "family": "Memory optimized",
    "enhanced_networking": false,
    "vCPU": 8,
    "generation": "previous",
    "ebs_iops": 125.0,
    "network_performance": "High",
    "ebs_throughput": 8000.0,
    "pretty_name": "M2 High Memory Quadruple Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.245",
            "yrTerm3.allUpfront": "0.7343",
            "yrTerm1.allUpfront": "1.0425",
            "yrTerm1.partialUpfront": "1.0637",
            "yrTerm3.partialUpfront": "0.7811"
          },
          "ondemand": "1.7"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.887",
            "yrTerm3.allUpfront": "1.2592",
            "yrTerm1.allUpfront": "1.5809",
            "yrTerm1.partialUpfront": "1.6132",
            "yrTerm3.partialUpfront": "1.3396"
          },
          "ondemand": "2.976"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.991",
            "yrTerm3.allUpfront": "0.5998",
            "yrTerm1.allUpfront": "0.8301",
            "yrTerm1.partialUpfront": "0.8471",
            "yrTerm3.partialUpfront": "0.6381"
          },
          "ondemand": "1.38"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.444",
            "yrTerm3.allUpfront": "0.2222",
            "yrTerm1.allUpfront": "0.3716",
            "yrTerm1.partialUpfront": "0.3792",
            "yrTerm3.partialUpfront": "0.2364"
          },
          "ondemand": "0.98"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.588",
            "yrTerm3.allUpfront": "0.9627",
            "yrTerm1.allUpfront": "1.3298",
            "yrTerm1.partialUpfront": "1.3569",
            "yrTerm3.partialUpfront": "1.0242"
          },
          "ondemand": "1.7"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.285",
            "yrTerm3.allUpfront": "1.5337",
            "yrTerm1.allUpfront": "1.9138",
            "yrTerm1.partialUpfront": "1.9528",
            "yrTerm3.partialUpfront": "1.6316"
          },
          "ondemand": "3.116"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.248",
            "yrTerm3.allUpfront": "0.7739",
            "yrTerm1.allUpfront": "1.0451",
            "yrTerm1.partialUpfront": "1.0665",
            "yrTerm3.partialUpfront": "0.8232"
          },
          "ondemand": "1.41"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.604",
            "yrTerm3.allUpfront": "0.3253",
            "yrTerm1.allUpfront": "0.5073",
            "yrTerm1.partialUpfront": "0.5177",
            "yrTerm3.partialUpfront": "0.3461"
          },
          "ondemand": "1.15"
        }
      },
      "sa-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.984",
            "yrTerm3.allUpfront": "1.0895",
            "yrTerm1.allUpfront": "1.6616",
            "yrTerm1.partialUpfront": "1.6955",
            "yrTerm3.partialUpfront": "1.159"
          },
          "ondemand": "2.411"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.705",
            "yrTerm3.allUpfront": "1.658",
            "yrTerm1.allUpfront": "2.2654",
            "yrTerm1.partialUpfront": "2.3117",
            "yrTerm3.partialUpfront": "1.7638"
          },
          "ondemand": "3.804"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.618",
            "yrTerm3.allUpfront": "0.9292",
            "yrTerm1.allUpfront": "1.5498",
            "yrTerm1.partialUpfront": "1.5814",
            "yrTerm3.partialUpfront": "0.9885"
          },
          "ondemand": "1.691"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.516",
            "yrTerm3.allUpfront": "0.3092",
            "yrTerm1.allUpfront": "0.4354",
            "yrTerm1.partialUpfront": "0.4443",
            "yrTerm3.partialUpfront": "0.329"
          },
          "ondemand": "1.291"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.549",
            "yrTerm3.allUpfront": "0.9367",
            "yrTerm1.allUpfront": "1.2974",
            "yrTerm1.partialUpfront": "1.3238",
            "yrTerm3.partialUpfront": "0.9965"
          },
          "ondemand": "1.7"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.183",
            "yrTerm3.allUpfront": "1.4553",
            "yrTerm1.allUpfront": "1.8285",
            "yrTerm1.partialUpfront": "1.8659",
            "yrTerm3.partialUpfront": "1.5482"
          },
          "ondemand": "2.976"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.226",
            "yrTerm3.allUpfront": "0.7577",
            "yrTerm1.allUpfront": "1.0272",
            "yrTerm1.partialUpfront": "1.0481",
            "yrTerm3.partialUpfront": "0.8061"
          },
          "ondemand": "1.483"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.592",
            "yrTerm3.allUpfront": "0.32",
            "yrTerm1.allUpfront": "0.5006",
            "yrTerm1.partialUpfront": "0.5109",
            "yrTerm3.partialUpfront": "0.3404"
          },
          "ondemand": "1.183"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.549",
            "yrTerm3.allUpfront": "0.9367",
            "yrTerm1.allUpfront": "1.2974",
            "yrTerm1.partialUpfront": "1.3238",
            "yrTerm3.partialUpfront": "0.9965"
          },
          "ondemand": "1.7"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.183",
            "yrTerm3.allUpfront": "1.4553",
            "yrTerm1.allUpfront": "1.8285",
            "yrTerm1.partialUpfront": "1.8659",
            "yrTerm3.partialUpfront": "1.5482"
          },
          "ondemand": "2.976"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.226",
            "yrTerm3.allUpfront": "0.7577",
            "yrTerm1.allUpfront": "1.0272",
            "yrTerm1.partialUpfront": "1.0481",
            "yrTerm3.partialUpfront": "0.8061"
          },
          "ondemand": "1.483"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.592",
            "yrTerm3.allUpfront": "0.32",
            "yrTerm1.allUpfront": "0.5006",
            "yrTerm1.partialUpfront": "0.5109",
            "yrTerm3.partialUpfront": "0.3404"
          },
          "ondemand": "1.183"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.245",
            "yrTerm3.allUpfront": "0.7343",
            "yrTerm1.allUpfront": "1.0425",
            "yrTerm1.partialUpfront": "1.0637",
            "yrTerm3.partialUpfront": "0.7811"
          },
          "ondemand": "1.7"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.887",
            "yrTerm3.allUpfront": "1.2592",
            "yrTerm1.allUpfront": "1.5809",
            "yrTerm1.partialUpfront": "1.6132",
            "yrTerm3.partialUpfront": "1.3396"
          },
          "ondemand": "2.976"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.991",
            "yrTerm3.allUpfront": "0.5998",
            "yrTerm1.allUpfront": "0.8301",
            "yrTerm1.partialUpfront": "0.8471",
            "yrTerm3.partialUpfront": "0.6381"
          },
          "ondemand": "1.38"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.444",
            "yrTerm3.allUpfront": "0.2222",
            "yrTerm1.allUpfront": "0.3716",
            "yrTerm1.partialUpfront": "0.3792",
            "yrTerm3.partialUpfront": "0.2364"
          },
          "ondemand": "0.98"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.131",
            "yrTerm3.allUpfront": "0.6856",
            "yrTerm1.allUpfront": "0.8796",
            "yrTerm1.partialUpfront": "0.9665",
            "yrTerm3.partialUpfront": "0.7534"
          },
          "ondemand": "2.051"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.029",
            "yrTerm3.allUpfront": "1.3104",
            "yrTerm1.allUpfront": "1.5783",
            "yrTerm1.partialUpfront": "1.7345",
            "yrTerm3.partialUpfront": "1.44"
          },
          "ondemand": "3.408"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.08",
            "yrTerm3.allUpfront": "0.6307",
            "yrTerm1.allUpfront": "0.8377",
            "yrTerm1.partialUpfront": "0.9235",
            "yrTerm3.partialUpfront": "0.7095"
          },
          "ondemand": "1.571"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.544",
            "yrTerm3.allUpfront": "0.2707",
            "yrTerm1.allUpfront": "0.4354",
            "yrTerm1.partialUpfront": "0.4693",
            "yrTerm3.partialUpfront": "0.297"
          },
          "ondemand": "1.171"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.549",
            "yrTerm3.allUpfront": "0.9367",
            "yrTerm1.allUpfront": "1.2974",
            "yrTerm1.partialUpfront": "1.3238",
            "yrTerm3.partialUpfront": "0.9965"
          },
          "ondemand": "1.844"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.183",
            "yrTerm3.allUpfront": "1.4553",
            "yrTerm1.allUpfront": "1.8285",
            "yrTerm1.partialUpfront": "1.8659",
            "yrTerm3.partialUpfront": "1.5482"
          },
          "ondemand": "3.178"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.252",
            "yrTerm3.allUpfront": "0.7746",
            "yrTerm1.allUpfront": "1.0487",
            "yrTerm1.partialUpfront": "1.0701",
            "yrTerm3.partialUpfront": "0.8241"
          },
          "ondemand": "1.5"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.536",
            "yrTerm3.allUpfront": "0.288",
            "yrTerm1.allUpfront": "0.4519",
            "yrTerm1.partialUpfront": "0.4612",
            "yrTerm3.partialUpfront": "0.3064"
          },
          "ondemand": "1.1"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.245",
            "yrTerm3.allUpfront": "0.7343",
            "yrTerm1.allUpfront": "1.0425",
            "yrTerm1.partialUpfront": "1.0637",
            "yrTerm3.partialUpfront": "0.7811"
          },
          "ondemand": "1.7"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.887",
            "yrTerm3.allUpfront": "1.2592",
            "yrTerm1.allUpfront": "1.5809",
            "yrTerm1.partialUpfront": "1.6132",
            "yrTerm3.partialUpfront": "1.3396"
          },
          "ondemand": "2.976"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.252",
            "yrTerm3.allUpfront": "0.7728",
            "yrTerm1.allUpfront": "1.0487",
            "yrTerm1.partialUpfront": "1.0701",
            "yrTerm3.partialUpfront": "0.8221"
          },
          "ondemand": "1.38"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.536",
            "yrTerm3.allUpfront": "0.2899",
            "yrTerm1.allUpfront": "0.4519",
            "yrTerm1.partialUpfront": "0.4612",
            "yrTerm3.partialUpfront": "0.3084"
          },
          "ondemand": "1.1"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 30,
      "max_enis": 8
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "PV"
    ],
    "ebs_optimized": true,
    "storage": {
      "ssd": false,
      "devices": 2,
      "size": 840.0
    },
    "instance_type": "m2.4xlarge",
    "ECU": 26.0,
    "memory": 68.4,
    "ebs_max_bandwidth": 1000.0
  },
  {
    "family": "Memory optimized",
    "enhanced_networking": false,
    "vCPU": 32,
    "generation": "previous",
    "ebs_iops": 0,
    "network_performance": "10 Gigabit",
    "ebs_throughput": 0,
    "pretty_name": "High Memory Cluster Eight Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "2.327",
            "yrTerm3.allUpfront": "1.3619",
            "yrTerm1.allUpfront": "1.9492",
            "yrTerm1.partialUpfront": "1.9889",
            "yrTerm3.partialUpfront": "1.4488"
          },
          "ondemand": "4.044"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "6.077",
            "yrTerm3.allUpfront": "4.375",
            "yrTerm1.allUpfront": "5.0904",
            "yrTerm1.partialUpfront": "5.1943",
            "yrTerm3.partialUpfront": "4.6543"
          },
          "ondemand": "7.25"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.077",
            "yrTerm3.allUpfront": "1.1609",
            "yrTerm1.allUpfront": "1.7397",
            "yrTerm1.partialUpfront": "1.7752",
            "yrTerm3.partialUpfront": "1.235"
          },
          "ondemand": "3.831"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.689",
            "yrTerm3.allUpfront": "0.8497",
            "yrTerm1.allUpfront": "1.4153",
            "yrTerm1.partialUpfront": "1.4442",
            "yrTerm3.partialUpfront": "0.904"
          },
          "ondemand": "3.5"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "2.693",
            "yrTerm3.allUpfront": "1.7283",
            "yrTerm1.allUpfront": "2.2558",
            "yrTerm1.partialUpfront": "2.3018",
            "yrTerm3.partialUpfront": "1.8386"
          },
          "ondemand": "4.623"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "6.266",
            "yrTerm3.allUpfront": "4.5988",
            "yrTerm1.allUpfront": "5.2487",
            "yrTerm1.partialUpfront": "5.3558",
            "yrTerm3.partialUpfront": "4.8923"
          },
          "ondemand": "7.676"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.455",
            "yrTerm3.allUpfront": "1.5383",
            "yrTerm1.allUpfront": "2.0566",
            "yrTerm1.partialUpfront": "2.0986",
            "yrTerm3.partialUpfront": "1.6365"
          },
          "ondemand": "4.42"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.086",
            "yrTerm3.allUpfront": "1.2422",
            "yrTerm1.allUpfront": "1.7478",
            "yrTerm1.partialUpfront": "1.7836",
            "yrTerm3.partialUpfront": "1.3215"
          },
          "ondemand": "4.105"
        }
      },
      "sa-east-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "2.327",
            "yrTerm3.allUpfront": "1.3619",
            "yrTerm1.allUpfront": "1.9492",
            "yrTerm1.partialUpfront": "1.9889",
            "yrTerm3.partialUpfront": "1.4488"
          },
          "ondemand": "4.044"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "6.077",
            "yrTerm3.allUpfront": "4.375",
            "yrTerm1.allUpfront": "5.0904",
            "yrTerm1.partialUpfront": "5.1943",
            "yrTerm3.partialUpfront": "4.6543"
          },
          "ondemand": "7.25"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.077",
            "yrTerm3.allUpfront": "1.1609",
            "yrTerm1.allUpfront": "1.7397",
            "yrTerm1.partialUpfront": "1.7752",
            "yrTerm3.partialUpfront": "1.235"
          },
          "ondemand": "3.831"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.689",
            "yrTerm3.allUpfront": "0.8497",
            "yrTerm1.allUpfront": "1.4153",
            "yrTerm1.partialUpfront": "1.4442",
            "yrTerm3.partialUpfront": "0.904"
          },
          "ondemand": "3.5"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "2.909",
            "yrTerm3.allUpfront": "1.7021",
            "yrTerm1.allUpfront": "2.4362",
            "yrTerm1.partialUpfront": "2.4859",
            "yrTerm3.partialUpfront": "1.8108"
          },
          "ondemand": "4.044"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "6.746",
            "yrTerm3.allUpfront": "5.4691",
            "yrTerm1.allUpfront": "6.3634",
            "yrTerm1.partialUpfront": "6.4933",
            "yrTerm3.partialUpfront": "5.8183"
          },
          "ondemand": "7.25"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.596",
            "yrTerm3.allUpfront": "1.4514",
            "yrTerm1.allUpfront": "2.1748",
            "yrTerm1.partialUpfront": "2.2192",
            "yrTerm3.partialUpfront": "1.544"
          },
          "ondemand": "3.831"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.832",
            "yrTerm3.allUpfront": "1.0904",
            "yrTerm1.allUpfront": "1.5348",
            "yrTerm1.partialUpfront": "1.5662",
            "yrTerm3.partialUpfront": "1.16"
          },
          "ondemand": "3.75"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 30,
      "max_enis": 8
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [],
    "ebs_optimized": false,
    "storage": {
      "ssd": true,
      "devices": 2,
      "size": 120.0
    },
    "instance_type": "cr1.8xlarge",
    "ECU": 88.0,
    "memory": 244.0,
    "ebs_max_bandwidth": 0
  },
  {
    "family": "Storage optimized",
    "enhanced_networking": false,
    "vCPU": 16,
    "generation": "previous",
    "ebs_iops": 0,
    "network_performance": "10 Gigabit",
    "ebs_throughput": 0,
    "pretty_name": "HI1. High I/O Quadruple Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "2.312",
            "yrTerm3.allUpfront": "1.3366",
            "yrTerm1.allUpfront": "1.9365",
            "yrTerm1.partialUpfront": "1.9758",
            "yrTerm3.partialUpfront": "1.4219"
          },
          "ondemand": "3.66"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.071",
            "yrTerm3.allUpfront": "1.9472",
            "yrTerm1.allUpfront": "2.5724",
            "yrTerm1.partialUpfront": "2.6249",
            "yrTerm3.partialUpfront": "2.0715"
          },
          "ondemand": "4.607"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.26",
            "yrTerm3.allUpfront": "1.2963",
            "yrTerm1.allUpfront": "1.8934",
            "yrTerm1.partialUpfront": "1.9321",
            "yrTerm3.partialUpfront": "1.379"
          },
          "ondemand": "3.58"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.698",
            "yrTerm3.allUpfront": "0.8451",
            "yrTerm1.allUpfront": "1.4229",
            "yrTerm1.partialUpfront": "1.4521",
            "yrTerm3.partialUpfront": "0.899"
          },
          "ondemand": "3.1"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "2.825",
            "yrTerm3.allUpfront": "1.8784",
            "yrTerm1.allUpfront": "2.3663",
            "yrTerm1.partialUpfront": "2.4146",
            "yrTerm3.partialUpfront": "1.9983"
          },
          "ondemand": "4.084"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "6.041",
            "yrTerm3.allUpfront": "4.4613",
            "yrTerm1.allUpfront": "5.0599",
            "yrTerm1.partialUpfront": "5.1632",
            "yrTerm3.partialUpfront": "4.7461"
          },
          "ondemand": "6.829"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.358",
            "yrTerm3.allUpfront": "1.5043",
            "yrTerm1.allUpfront": "1.9755",
            "yrTerm1.partialUpfront": "2.0157",
            "yrTerm3.partialUpfront": "1.6003"
          },
          "ondemand": "3.686"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.877",
            "yrTerm3.allUpfront": "1.1179",
            "yrTerm1.allUpfront": "1.5726",
            "yrTerm1.partialUpfront": "1.6047",
            "yrTerm3.partialUpfront": "1.1893"
          },
          "ondemand": "3.276"
        }
      },
      "sa-east-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "2.312",
            "yrTerm3.allUpfront": "1.3366",
            "yrTerm1.allUpfront": "1.9365",
            "yrTerm1.partialUpfront": "1.9758",
            "yrTerm3.partialUpfront": "1.4219"
          },
          "ondemand": "3.66"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.071",
            "yrTerm3.allUpfront": "1.9472",
            "yrTerm1.allUpfront": "2.5724",
            "yrTerm1.partialUpfront": "2.6249",
            "yrTerm3.partialUpfront": "2.0715"
          },
          "ondemand": "4.607"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.26",
            "yrTerm3.allUpfront": "1.2963",
            "yrTerm1.allUpfront": "1.8934",
            "yrTerm1.partialUpfront": "1.9321",
            "yrTerm3.partialUpfront": "1.379"
          },
          "ondemand": "3.58"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.698",
            "yrTerm3.allUpfront": "0.8451",
            "yrTerm1.allUpfront": "1.4229",
            "yrTerm1.partialUpfront": "1.4521",
            "yrTerm3.partialUpfront": "0.899"
          },
          "ondemand": "3.1"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "2.312",
            "yrTerm3.allUpfront": "1.3366",
            "yrTerm1.allUpfront": "1.9365",
            "yrTerm1.partialUpfront": "1.9758",
            "yrTerm3.partialUpfront": "1.4219"
          },
          "ondemand": "3.66"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.071",
            "yrTerm3.allUpfront": "1.9472",
            "yrTerm1.allUpfront": "2.5724",
            "yrTerm1.partialUpfront": "2.6249",
            "yrTerm3.partialUpfront": "2.0715"
          },
          "ondemand": "4.607"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.26",
            "yrTerm3.allUpfront": "1.2963",
            "yrTerm1.allUpfront": "1.8934",
            "yrTerm1.partialUpfront": "1.9321",
            "yrTerm3.partialUpfront": "1.379"
          },
          "ondemand": "3.58"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.698",
            "yrTerm3.allUpfront": "0.8451",
            "yrTerm1.allUpfront": "1.4229",
            "yrTerm1.partialUpfront": "1.4521",
            "yrTerm3.partialUpfront": "0.899"
          },
          "ondemand": "3.1"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 30,
      "max_enis": 8
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "HVM",
      "PV"
    ],
    "ebs_optimized": false,
    "storage": {
      "ssd": true,
      "devices": 2,
      "size": 1024.0
    },
    "instance_type": "hi1.4xlarge",
    "ECU": 35.0,
    "memory": 60.5,
    "ebs_max_bandwidth": 0
  },
  {
    "family": "Storage optimized",
    "enhanced_networking": false,
    "vCPU": 16,
    "generation": "previous",
    "ebs_iops": 0,
    "network_performance": "10 Gigabit",
    "ebs_throughput": 0,
    "pretty_name": "High Storage Eight Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.211",
            "yrTerm3.allUpfront": "1.8317",
            "yrTerm1.allUpfront": "2.6898",
            "yrTerm1.partialUpfront": "2.7447",
            "yrTerm3.partialUpfront": "1.9486"
          },
          "ondemand": "5.144"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "6.961",
            "yrTerm3.allUpfront": "4.845",
            "yrTerm1.allUpfront": "5.8307",
            "yrTerm1.partialUpfront": "5.9496",
            "yrTerm3.partialUpfront": "5.1542"
          },
          "ondemand": "8.35"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.961",
            "yrTerm3.allUpfront": "1.6309",
            "yrTerm1.allUpfront": "2.4804",
            "yrTerm1.partialUpfront": "2.531",
            "yrTerm3.partialUpfront": "1.735"
          },
          "ondemand": "4.931"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.574",
            "yrTerm3.allUpfront": "1.3197",
            "yrTerm1.allUpfront": "2.1559",
            "yrTerm1.partialUpfront": "2.2",
            "yrTerm3.partialUpfront": "1.404"
          },
          "ondemand": "4.6"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.859",
            "yrTerm3.allUpfront": "2.4236",
            "yrTerm1.allUpfront": "3.2326",
            "yrTerm1.partialUpfront": "3.2986",
            "yrTerm3.partialUpfront": "2.5783"
          },
          "ondemand": "5.917"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "7.432",
            "yrTerm3.allUpfront": "5.2932",
            "yrTerm1.allUpfront": "6.2255",
            "yrTerm1.partialUpfront": "6.3526",
            "yrTerm3.partialUpfront": "5.631"
          },
          "ondemand": "8.97"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "3.622",
            "yrTerm3.allUpfront": "2.2327",
            "yrTerm1.allUpfront": "3.0338",
            "yrTerm1.partialUpfront": "3.0957",
            "yrTerm3.partialUpfront": "2.3752"
          },
          "ondemand": "5.714"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "3.253",
            "yrTerm3.allUpfront": "1.9366",
            "yrTerm1.allUpfront": "2.725",
            "yrTerm1.partialUpfront": "2.7807",
            "yrTerm3.partialUpfront": "2.0602"
          },
          "ondemand": "5.4"
        }
      },
      "sa-east-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "4.019",
            "yrTerm3.allUpfront": "2.3914",
            "yrTerm1.allUpfront": "3.3659",
            "yrTerm1.partialUpfront": "3.4346",
            "yrTerm3.partialUpfront": "2.544"
          },
          "ondemand": "6.114"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "9.082",
            "yrTerm3.allUpfront": "5.405",
            "yrTerm1.allUpfront": "7.6073",
            "yrTerm1.partialUpfront": "7.7625",
            "yrTerm3.partialUpfront": "5.75"
          },
          "ondemand": "9.32"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "3.682",
            "yrTerm3.allUpfront": "2.1916",
            "yrTerm1.allUpfront": "3.0841",
            "yrTerm1.partialUpfront": "3.1471",
            "yrTerm3.partialUpfront": "2.3315"
          },
          "ondemand": "5.901"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "3.158",
            "yrTerm3.allUpfront": "1.88",
            "yrTerm1.allUpfront": "2.6459",
            "yrTerm1.partialUpfront": "2.7",
            "yrTerm3.partialUpfront": "2"
          },
          "ondemand": "5.57"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "4.019",
            "yrTerm3.allUpfront": "2.3914",
            "yrTerm1.allUpfront": "3.3659",
            "yrTerm1.partialUpfront": "3.4346",
            "yrTerm3.partialUpfront": "2.544"
          },
          "ondemand": "6.114"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "9.082",
            "yrTerm3.allUpfront": "5.405",
            "yrTerm1.allUpfront": "7.6073",
            "yrTerm1.partialUpfront": "7.7625",
            "yrTerm3.partialUpfront": "5.75"
          },
          "ondemand": "9.32"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "3.682",
            "yrTerm3.allUpfront": "2.1916",
            "yrTerm1.allUpfront": "3.0841",
            "yrTerm1.partialUpfront": "3.1471",
            "yrTerm3.partialUpfront": "2.3315"
          },
          "ondemand": "5.901"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "3.158",
            "yrTerm3.allUpfront": "1.88",
            "yrTerm1.allUpfront": "2.6459",
            "yrTerm1.partialUpfront": "2.7",
            "yrTerm3.partialUpfront": "2"
          },
          "ondemand": "5.57"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.211",
            "yrTerm3.allUpfront": "1.8317",
            "yrTerm1.allUpfront": "2.6898",
            "yrTerm1.partialUpfront": "2.7447",
            "yrTerm3.partialUpfront": "1.9486"
          },
          "ondemand": "5.144"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "6.961",
            "yrTerm3.allUpfront": "4.845",
            "yrTerm1.allUpfront": "5.8307",
            "yrTerm1.partialUpfront": "5.9496",
            "yrTerm3.partialUpfront": "5.1542"
          },
          "ondemand": "8.35"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.961",
            "yrTerm3.allUpfront": "1.6309",
            "yrTerm1.allUpfront": "2.4804",
            "yrTerm1.partialUpfront": "2.531",
            "yrTerm3.partialUpfront": "1.735"
          },
          "ondemand": "4.931"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.574",
            "yrTerm3.allUpfront": "1.3197",
            "yrTerm1.allUpfront": "2.1559",
            "yrTerm1.partialUpfront": "2.2",
            "yrTerm3.partialUpfront": "1.404"
          },
          "ondemand": "4.6"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.727",
            "yrTerm3.allUpfront": "2.0964",
            "yrTerm1.allUpfront": "3.1218",
            "yrTerm1.partialUpfront": "3.1856",
            "yrTerm3.partialUpfront": "2.2302"
          },
          "ondemand": "6.339"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "7.476",
            "yrTerm3.allUpfront": "5.1097",
            "yrTerm1.allUpfront": "6.2619",
            "yrTerm1.partialUpfront": "6.3897",
            "yrTerm3.partialUpfront": "5.4359"
          },
          "ondemand": "9.27"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "3.477",
            "yrTerm3.allUpfront": "1.8956",
            "yrTerm1.allUpfront": "2.9122",
            "yrTerm1.partialUpfront": "2.9716",
            "yrTerm3.partialUpfront": "2.0166"
          },
          "ondemand": "6.126"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "3.088",
            "yrTerm3.allUpfront": "1.5837",
            "yrTerm1.allUpfront": "2.5872",
            "yrTerm1.partialUpfront": "2.64",
            "yrTerm3.partialUpfront": "1.6848"
          },
          "ondemand": "5.52"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "4.014",
            "yrTerm3.allUpfront": "2.2895",
            "yrTerm1.allUpfront": "3.3621",
            "yrTerm1.partialUpfront": "3.4307",
            "yrTerm3.partialUpfront": "2.4356"
          },
          "ondemand": "5.144"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "7.742",
            "yrTerm3.allUpfront": "6.0557",
            "yrTerm1.allUpfront": "7.2889",
            "yrTerm1.partialUpfront": "7.4376",
            "yrTerm3.partialUpfront": "6.4422"
          },
          "ondemand": "8.35"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "3.701",
            "yrTerm3.allUpfront": "2.0379",
            "yrTerm1.allUpfront": "3.0998",
            "yrTerm1.partialUpfront": "3.163",
            "yrTerm3.partialUpfront": "2.168"
          },
          "ondemand": "4.931"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.858",
            "yrTerm3.allUpfront": "1.7014",
            "yrTerm1.allUpfront": "2.3941",
            "yrTerm1.partialUpfront": "2.443",
            "yrTerm3.partialUpfront": "1.81"
          },
          "ondemand": "4.9"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 30,
      "max_enis": 8
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "HVM",
      "PV"
    ],
    "ebs_optimized": false,
    "storage": {
      "ssd": false,
      "devices": 24,
      "size": 2000.0
    },
    "instance_type": "hs1.8xlarge",
    "ECU": 35.0,
    "memory": 117.0,
    "ebs_max_bandwidth": 0
  },
  {
    "family": "Micro",
    "enhanced_networking": false,
    "vCPU": 1,
    "generation": "previous",
    "ebs_iops": 0,
    "network_performance": "Very Low",
    "ebs_throughput": 0,
    "pretty_name": "T1 Micro",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.047",
            "yrTerm3.allUpfront": "0.0343",
            "yrTerm1.allUpfront": "0.0392",
            "yrTerm1.partialUpfront": "0.04",
            "yrTerm3.partialUpfront": "0.0365"
          },
          "ondemand": "0.07"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.015",
            "yrTerm3.allUpfront": "0.0102",
            "yrTerm1.allUpfront": "0.0128",
            "yrTerm1.partialUpfront": "0.0131",
            "yrTerm3.partialUpfront": "0.0108"
          },
          "ondemand": "0.02"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.014",
            "yrTerm3.allUpfront": "0.0083",
            "yrTerm1.allUpfront": "0.0118",
            "yrTerm1.partialUpfront": "0.0121",
            "yrTerm3.partialUpfront": "0.0088"
          },
          "ondemand": "0.02"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.058",
            "yrTerm3.allUpfront": "0.0433",
            "yrTerm1.allUpfront": "0.0485",
            "yrTerm1.partialUpfront": "0.0495",
            "yrTerm3.partialUpfront": "0.0461"
          },
          "ondemand": "0.077"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.029",
            "yrTerm3.allUpfront": "0.0205",
            "yrTerm1.allUpfront": "0.0245",
            "yrTerm1.partialUpfront": "0.0251",
            "yrTerm3.partialUpfront": "0.0218"
          },
          "ondemand": "0.033"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.018",
            "yrTerm3.allUpfront": "0.012",
            "yrTerm1.allUpfront": "0.0158",
            "yrTerm1.partialUpfront": "0.0161",
            "yrTerm3.partialUpfront": "0.0128"
          },
          "ondemand": "0.026"
        }
      },
      "sa-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.048",
            "yrTerm3.allUpfront": "0.0357",
            "yrTerm1.allUpfront": "0.0403",
            "yrTerm1.partialUpfront": "0.0411",
            "yrTerm3.partialUpfront": "0.038"
          },
          "ondemand": "0.077"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.026",
            "yrTerm3.allUpfront": "0.017",
            "yrTerm1.allUpfront": "0.0221",
            "yrTerm1.partialUpfront": "0.0226",
            "yrTerm3.partialUpfront": "0.0181"
          },
          "ondemand": "0.037"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.019",
            "yrTerm3.allUpfront": "0.0114",
            "yrTerm1.allUpfront": "0.0162",
            "yrTerm1.partialUpfront": "0.0166",
            "yrTerm3.partialUpfront": "0.0121"
          },
          "ondemand": "0.027"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.05",
            "yrTerm3.allUpfront": "0.0372",
            "yrTerm1.allUpfront": "0.0422",
            "yrTerm1.partialUpfront": "0.0431",
            "yrTerm3.partialUpfront": "0.0396"
          },
          "ondemand": "0.075"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.019",
            "yrTerm3.allUpfront": "0.012",
            "yrTerm1.allUpfront": "0.0158",
            "yrTerm1.partialUpfront": "0.0161",
            "yrTerm3.partialUpfront": "0.0128"
          },
          "ondemand": "0.02"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.016",
            "yrTerm3.allUpfront": "0.0111",
            "yrTerm1.allUpfront": "0.0147",
            "yrTerm1.partialUpfront": "0.0151",
            "yrTerm3.partialUpfront": "0.0118"
          },
          "ondemand": "0.02"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.05",
            "yrTerm3.allUpfront": "0.0372",
            "yrTerm1.allUpfront": "0.0422",
            "yrTerm1.partialUpfront": "0.0431",
            "yrTerm3.partialUpfront": "0.0396"
          },
          "ondemand": "0.075"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.019",
            "yrTerm3.allUpfront": "0.012",
            "yrTerm1.allUpfront": "0.0158",
            "yrTerm1.partialUpfront": "0.0161",
            "yrTerm3.partialUpfront": "0.0128"
          },
          "ondemand": "0.02"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.016",
            "yrTerm3.allUpfront": "0.0111",
            "yrTerm1.allUpfront": "0.0147",
            "yrTerm1.partialUpfront": "0.0151",
            "yrTerm3.partialUpfront": "0.0118"
          },
          "ondemand": "0.02"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.047",
            "yrTerm3.allUpfront": "0.0343",
            "yrTerm1.allUpfront": "0.0392",
            "yrTerm1.partialUpfront": "0.04",
            "yrTerm3.partialUpfront": "0.0365"
          },
          "ondemand": "0.07"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.015",
            "yrTerm3.allUpfront": "0.0102",
            "yrTerm1.allUpfront": "0.0128",
            "yrTerm1.partialUpfront": "0.0131",
            "yrTerm3.partialUpfront": "0.0108"
          },
          "ondemand": "0.02"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.014",
            "yrTerm3.allUpfront": "0.0083",
            "yrTerm1.allUpfront": "0.0118",
            "yrTerm1.partialUpfront": "0.0121",
            "yrTerm3.partialUpfront": "0.0088"
          },
          "ondemand": "0.02"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.066",
            "yrTerm3.allUpfront": "0.0494",
            "yrTerm1.allUpfront": "0.0511",
            "yrTerm1.partialUpfront": "0.0568",
            "yrTerm3.partialUpfront": "0.0543"
          },
          "ondemand": "0.074"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.022",
            "yrTerm3.allUpfront": "0.014",
            "yrTerm1.allUpfront": "0.0172",
            "yrTerm1.partialUpfront": "0.0191",
            "yrTerm3.partialUpfront": "0.0158"
          },
          "ondemand": "0.025"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.015",
            "yrTerm3.allUpfront": "0.0087",
            "yrTerm1.allUpfront": "0.0119",
            "yrTerm1.partialUpfront": "0.0131",
            "yrTerm3.partialUpfront": "0.0098"
          },
          "ondemand": "0.024"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.05",
            "yrTerm3.allUpfront": "0.0372",
            "yrTerm1.allUpfront": "0.0422",
            "yrTerm1.partialUpfront": "0.0431",
            "yrTerm3.partialUpfront": "0.0396"
          },
          "ondemand": "0.075"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.025",
            "yrTerm3.allUpfront": "0.0167",
            "yrTerm1.allUpfront": "0.0207",
            "yrTerm1.partialUpfront": "0.0211",
            "yrTerm3.partialUpfront": "0.0178"
          },
          "ondemand": "0.035"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.017",
            "yrTerm3.allUpfront": "0.0111",
            "yrTerm1.allUpfront": "0.0147",
            "yrTerm1.partialUpfront": "0.0151",
            "yrTerm3.partialUpfront": "0.0118"
          },
          "ondemand": "0.025"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "ondemand": "N/A"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "ondemand": "N/A"
        },
        "linux": {
          "ondemand": "N/A"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.047",
            "yrTerm3.allUpfront": "0.0343",
            "yrTerm1.allUpfront": "0.0392",
            "yrTerm1.partialUpfront": "0.04",
            "yrTerm3.partialUpfront": "0.0365"
          },
          "ondemand": "0.075"
        },
        "mswinSQL": {
          "ondemand": "N/A"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.019",
            "yrTerm3.allUpfront": "0.012",
            "yrTerm1.allUpfront": "0.0158",
            "yrTerm1.partialUpfront": "0.0161",
            "yrTerm3.partialUpfront": "0.0128"
          },
          "ondemand": "0.02"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.016",
            "yrTerm3.allUpfront": "0.0111",
            "yrTerm1.allUpfront": "0.0147",
            "yrTerm1.partialUpfront": "0.0151",
            "yrTerm3.partialUpfront": "0.0118"
          },
          "ondemand": "0.02"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 2,
      "max_enis": 2
    },
    "arch": [
      "i386",
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "PV"
    ],
    "ebs_optimized": false,
    "storage": null,
    "instance_type": "t1.micro",
    "ECU": "variable",
    "memory": 0.613,
    "ebs_max_bandwidth": 0
  },
  {
    "family": "General purpose",
    "enhanced_networking": false,
    "vCPU": 1,
    "generation": "current",
    "ebs_iops": 0,
    "network_performance": "Low",
    "ebs_throughput": 0,
    "pretty_name": "T2 Nano",
    "pricing": {
      "us-east-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.0068",
            "yrTerm3.allUpfront": "0.0052",
            "yrTerm1.allUpfront": "0.0067",
            "yrTerm1.partialUpfront": "0.0067",
            "yrTerm3.partialUpfront": "0.0054"
          },
          "ondemand": "0.0088"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.0045",
            "yrTerm3.allUpfront": "0.0029",
            "yrTerm1.allUpfront": "0.0043",
            "yrTerm1.partialUpfront": "0.0044",
            "yrTerm3.partialUpfront": "0.0031"
          },
          "ondemand": "0.0065"
        }
      },
      "ap-northeast-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.0098",
            "yrTerm3.allUpfront": "0.007",
            "yrTerm1.allUpfront": "0.0096",
            "yrTerm1.partialUpfront": "0.0097",
            "yrTerm3.partialUpfront": "0.0073"
          },
          "ondemand": "0.0123"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.0075",
            "yrTerm3.allUpfront": "0.0047",
            "yrTerm1.allUpfront": "0.0072",
            "yrTerm1.partialUpfront": "0.0074",
            "yrTerm3.partialUpfront": "0.005"
          },
          "ondemand": "0.01"
        }
      },
      "sa-east-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.0113",
            "yrTerm3.allUpfront": "0.008",
            "yrTerm1.allUpfront": "0.0102",
            "yrTerm1.partialUpfront": "0.0102",
            "yrTerm3.partialUpfront": "0.0084"
          },
          "ondemand": "0.0158"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.009",
            "yrTerm3.allUpfront": "0.0057",
            "yrTerm1.allUpfront": "0.0078",
            "yrTerm1.partialUpfront": "0.0079",
            "yrTerm3.partialUpfront": "0.0061"
          },
          "ondemand": "0.0135"
        }
      },
      "ap-northeast-2": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.0109",
            "yrTerm3.allUpfront": "0.007",
            "yrTerm1.allUpfront": "0.0095",
            "yrTerm1.partialUpfront": "0.0097",
            "yrTerm3.partialUpfront": "0.0073"
          },
          "ondemand": "0.012"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.0086",
            "yrTerm3.allUpfront": "0.0047",
            "yrTerm1.allUpfront": "0.0072",
            "yrTerm1.partialUpfront": "0.0074",
            "yrTerm3.partialUpfront": "0.005"
          },
          "ondemand": "0.01"
        }
      },
      "ap-southeast-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.0103",
            "yrTerm3.allUpfront": "0.0078",
            "yrTerm1.allUpfront": "0.0097",
            "yrTerm1.partialUpfront": "0.0098",
            "yrTerm3.partialUpfront": "0.0082"
          },
          "ondemand": "0.0123"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.008",
            "yrTerm3.allUpfront": "0.0055",
            "yrTerm1.allUpfront": "0.0073",
            "yrTerm1.partialUpfront": "0.0075",
            "yrTerm3.partialUpfront": "0.0059"
          },
          "ondemand": "0.01"
        }
      },
      "ap-southeast-2": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.0103",
            "yrTerm3.allUpfront": "0.0078",
            "yrTerm1.allUpfront": "0.0097",
            "yrTerm1.partialUpfront": "0.0099",
            "yrTerm3.partialUpfront": "0.0082"
          },
          "ondemand": "0.0123"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.008",
            "yrTerm3.allUpfront": "0.0055",
            "yrTerm1.allUpfront": "0.0073",
            "yrTerm1.partialUpfront": "0.0076",
            "yrTerm3.partialUpfront": "0.0059"
          },
          "ondemand": "0.01"
        }
      },
      "us-west-2": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.0068",
            "yrTerm3.allUpfront": "0.0052",
            "yrTerm1.allUpfront": "0.0067",
            "yrTerm1.partialUpfront": "0.0067",
            "yrTerm3.partialUpfront": "0.0054"
          },
          "ondemand": "0.0088"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.0045",
            "yrTerm3.allUpfront": "0.0029",
            "yrTerm1.allUpfront": "0.0043",
            "yrTerm1.partialUpfront": "0.0044",
            "yrTerm3.partialUpfront": "0.0031"
          },
          "ondemand": "0.0065"
        }
      },
      "us-gov-west-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.0078",
            "yrTerm3.allUpfront": "0.0056",
            "yrTerm1.allUpfront": "0.0073",
            "yrTerm1.partialUpfront": "0.0074",
            "yrTerm3.partialUpfront": "0.0058"
          },
          "ondemand": "0.0098"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.0055",
            "yrTerm3.allUpfront": "0.0033",
            "yrTerm1.allUpfront": "0.0049",
            "yrTerm1.partialUpfront": "0.005",
            "yrTerm3.partialUpfront": "0.0035"
          },
          "ondemand": "0.0075"
        }
      },
      "us-west-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.0083",
            "yrTerm3.allUpfront": "0.0063",
            "yrTerm1.allUpfront": "0.0081",
            "yrTerm1.partialUpfront": "0.0082",
            "yrTerm3.partialUpfront": "0.0065"
          },
          "ondemand": "0.0108"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.006",
            "yrTerm3.allUpfront": "0.004",
            "yrTerm1.allUpfront": "0.0057",
            "yrTerm1.partialUpfront": "0.0059",
            "yrTerm3.partialUpfront": "0.0042"
          },
          "ondemand": "0.0085"
        }
      },
      "eu-central-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.0078",
            "yrTerm3.allUpfront": "0.0057",
            "yrTerm1.allUpfront": "0.0074",
            "yrTerm1.partialUpfront": "0.0075",
            "yrTerm3.partialUpfront": "0.0059"
          },
          "ondemand": "0.0098"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.0055",
            "yrTerm3.allUpfront": "0.0034",
            "yrTerm1.allUpfront": "0.005",
            "yrTerm1.partialUpfront": "0.0052",
            "yrTerm3.partialUpfront": "0.0036"
          },
          "ondemand": "0.0075"
        }
      },
      "eu-west-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.0073",
            "yrTerm3.allUpfront": "0.0055",
            "yrTerm1.allUpfront": "0.0071",
            "yrTerm1.partialUpfront": "0.0071",
            "yrTerm3.partialUpfront": "0.0057"
          },
          "ondemand": "0.0093"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.005",
            "yrTerm3.allUpfront": "0.0032",
            "yrTerm1.allUpfront": "0.0047",
            "yrTerm1.partialUpfront": "0.0048",
            "yrTerm3.partialUpfront": "0.0034"
          },
          "ondemand": "0.007"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 2,
      "max_enis": 2
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": true,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": false,
    "storage": null,
    "instance_type": "t2.nano",
    "ECU": "variable",
    "memory": 0.5,
    "ebs_max_bandwidth": 0
  },
  {
    "family": "General purpose",
    "enhanced_networking": false,
    "vCPU": 1,
    "generation": "current",
    "ebs_iops": 0,
    "network_performance": "Low to Moderate",
    "ebs_throughput": 0,
    "pretty_name": "T2 Micro",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.047",
            "yrTerm3.allUpfront": "0.0343",
            "yrTerm1.allUpfront": "0.039",
            "yrTerm1.partialUpfront": "0.0399",
            "yrTerm3.partialUpfront": "0.0365"
          },
          "ondemand": "0.068"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.014",
            "yrTerm3.allUpfront": "0.0089",
            "yrTerm1.allUpfront": "0.012",
            "yrTerm1.partialUpfront": "0.0122",
            "yrTerm3.partialUpfront": "0.0095"
          },
          "ondemand": "0.018"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.009",
            "yrTerm3.allUpfront": "0.0057",
            "yrTerm1.allUpfront": "0.0086",
            "yrTerm1.partialUpfront": "0.0088",
            "yrTerm3.partialUpfront": "0.0061"
          },
          "ondemand": "0.013"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.052",
            "yrTerm3.allUpfront": "0.0376",
            "yrTerm1.allUpfront": "0.0438",
            "yrTerm1.partialUpfront": "0.0447",
            "yrTerm3.partialUpfront": "0.04"
          },
          "ondemand": "0.075"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.02",
            "yrTerm3.allUpfront": "0.0121",
            "yrTerm1.allUpfront": "0.0171",
            "yrTerm1.partialUpfront": "0.0175",
            "yrTerm3.partialUpfront": "0.0128"
          },
          "ondemand": "0.025"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.015",
            "yrTerm3.allUpfront": "0.0094",
            "yrTerm1.allUpfront": "0.0144",
            "yrTerm1.partialUpfront": "0.0147",
            "yrTerm3.partialUpfront": "0.01"
          },
          "ondemand": "0.02"
        }
      },
      "sa-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.048",
            "yrTerm3.allUpfront": "0.036",
            "yrTerm1.allUpfront": "0.04",
            "yrTerm1.partialUpfront": "0.0408",
            "yrTerm3.partialUpfront": "0.0383"
          },
          "ondemand": "0.077"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.023",
            "yrTerm3.allUpfront": "0.0142",
            "yrTerm1.allUpfront": "0.0192",
            "yrTerm1.partialUpfront": "0.0196",
            "yrTerm3.partialUpfront": "0.0151"
          },
          "ondemand": "0.032"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.018",
            "yrTerm3.allUpfront": "0.0114",
            "yrTerm1.allUpfront": "0.0154",
            "yrTerm1.partialUpfront": "0.0158",
            "yrTerm3.partialUpfront": "0.0121"
          },
          "ondemand": "0.027"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.0722",
            "yrTerm3.allUpfront": "0.0644",
            "yrTerm1.allUpfront": "0.0694",
            "yrTerm1.partialUpfront": "0.0697",
            "yrTerm3.partialUpfront": "0.065"
          },
          "ondemand": "0.075"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.0218",
            "yrTerm3.allUpfront": "0.014",
            "yrTerm1.allUpfront": "0.019",
            "yrTerm1.partialUpfront": "0.0193",
            "yrTerm3.partialUpfront": "0.0146"
          },
          "ondemand": "0.025"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.0172",
            "yrTerm3.allUpfront": "0.0094",
            "yrTerm1.allUpfront": "0.0144",
            "yrTerm1.partialUpfront": "0.0147",
            "yrTerm3.partialUpfront": "0.01"
          },
          "ondemand": "0.02"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.05",
            "yrTerm3.allUpfront": "0.0373",
            "yrTerm1.allUpfront": "0.0422",
            "yrTerm1.partialUpfront": "0.0431",
            "yrTerm3.partialUpfront": "0.0397"
          },
          "ondemand": "0.075"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.019",
            "yrTerm3.allUpfront": "0.0119",
            "yrTerm1.allUpfront": "0.0162",
            "yrTerm1.partialUpfront": "0.0165",
            "yrTerm3.partialUpfront": "0.0127"
          },
          "ondemand": "0.02"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.016",
            "yrTerm3.allUpfront": "0.011",
            "yrTerm1.allUpfront": "0.0146",
            "yrTerm1.partialUpfront": "0.015",
            "yrTerm3.partialUpfront": "0.0118"
          },
          "ondemand": "0.02"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.05",
            "yrTerm3.allUpfront": "0.0373",
            "yrTerm1.allUpfront": "0.0422",
            "yrTerm1.partialUpfront": "0.0431",
            "yrTerm3.partialUpfront": "0.0397"
          },
          "ondemand": "0.075"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.019",
            "yrTerm3.allUpfront": "0.0119",
            "yrTerm1.allUpfront": "0.0162",
            "yrTerm1.partialUpfront": "0.0165",
            "yrTerm3.partialUpfront": "0.0127"
          },
          "ondemand": "0.02"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.016",
            "yrTerm3.allUpfront": "0.011",
            "yrTerm1.allUpfront": "0.0146",
            "yrTerm1.partialUpfront": "0.015",
            "yrTerm3.partialUpfront": "0.0118"
          },
          "ondemand": "0.02"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.047",
            "yrTerm3.allUpfront": "0.0343",
            "yrTerm1.allUpfront": "0.039",
            "yrTerm1.partialUpfront": "0.0399",
            "yrTerm3.partialUpfront": "0.0365"
          },
          "ondemand": "0.068"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.014",
            "yrTerm3.allUpfront": "0.0089",
            "yrTerm1.allUpfront": "0.012",
            "yrTerm1.partialUpfront": "0.0122",
            "yrTerm3.partialUpfront": "0.0095"
          },
          "ondemand": "0.018"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.009",
            "yrTerm3.allUpfront": "0.0057",
            "yrTerm1.allUpfront": "0.0086",
            "yrTerm1.partialUpfront": "0.0088",
            "yrTerm3.partialUpfront": "0.0061"
          },
          "ondemand": "0.013"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.049",
            "yrTerm3.allUpfront": "0.036",
            "yrTerm1.allUpfront": "0.041",
            "yrTerm1.partialUpfront": "0.0419",
            "yrTerm3.partialUpfront": "0.0383"
          },
          "ondemand": "0.071"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.017",
            "yrTerm3.allUpfront": "0.0105",
            "yrTerm1.allUpfront": "0.0139",
            "yrTerm1.partialUpfront": "0.0142",
            "yrTerm3.partialUpfront": "0.0112"
          },
          "ondemand": "0.021"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.011",
            "yrTerm3.allUpfront": "0.0065",
            "yrTerm1.allUpfront": "0.0097",
            "yrTerm1.partialUpfront": "0.01",
            "yrTerm3.partialUpfront": "0.0069"
          },
          "ondemand": "0.015"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.05",
            "yrTerm3.allUpfront": "0.0358",
            "yrTerm1.allUpfront": "0.0416",
            "yrTerm1.partialUpfront": "0.0425",
            "yrTerm3.partialUpfront": "0.038"
          },
          "ondemand": "0.072"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.02",
            "yrTerm3.allUpfront": "0.0112",
            "yrTerm1.allUpfront": "0.0168",
            "yrTerm1.partialUpfront": "0.0172",
            "yrTerm3.partialUpfront": "0.0119"
          },
          "ondemand": "0.022"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.012",
            "yrTerm3.allUpfront": "0.008",
            "yrTerm1.allUpfront": "0.0114",
            "yrTerm1.partialUpfront": "0.0118",
            "yrTerm3.partialUpfront": "0.0085"
          },
          "ondemand": "0.017"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.049",
            "yrTerm3.allUpfront": "0.035",
            "yrTerm1.allUpfront": "0.0409",
            "yrTerm1.partialUpfront": "0.0417",
            "yrTerm3.partialUpfront": "0.0372"
          },
          "ondemand": "0.07"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.016",
            "yrTerm3.allUpfront": "0.0095",
            "yrTerm1.allUpfront": "0.0136",
            "yrTerm1.partialUpfront": "0.0139",
            "yrTerm3.partialUpfront": "0.0101"
          },
          "ondemand": "0.02"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.011",
            "yrTerm3.allUpfront": "0.0068",
            "yrTerm1.allUpfront": "0.01",
            "yrTerm1.partialUpfront": "0.0103",
            "yrTerm3.partialUpfront": "0.0073"
          },
          "ondemand": "0.015"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.047",
            "yrTerm3.allUpfront": "0.0345",
            "yrTerm1.allUpfront": "0.0392",
            "yrTerm1.partialUpfront": "0.04",
            "yrTerm3.partialUpfront": "0.0367"
          },
          "ondemand": "0.069"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.015",
            "yrTerm3.allUpfront": "0.0092",
            "yrTerm1.allUpfront": "0.0124",
            "yrTerm1.partialUpfront": "0.0126",
            "yrTerm3.partialUpfront": "0.0098"
          },
          "ondemand": "0.019"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.01",
            "yrTerm3.allUpfront": "0.0063",
            "yrTerm1.allUpfront": "0.0092",
            "yrTerm1.partialUpfront": "0.0095",
            "yrTerm3.partialUpfront": "0.0067"
          },
          "ondemand": "0.014"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 2,
      "max_enis": 2
    },
    "arch": [
      "x86_64",
      "i386"
    ],
    "vpc_only": true,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": false,
    "storage": null,
    "instance_type": "t2.micro",
    "ECU": "variable",
    "memory": 1.0,
    "ebs_max_bandwidth": 0
  },
  {
    "family": "General purpose",
    "enhanced_networking": false,
    "vCPU": 1,
    "generation": "current",
    "ebs_iops": 0,
    "network_performance": "Low to Moderate",
    "ebs_throughput": 0,
    "pretty_name": "T2 Small",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.093",
            "yrTerm3.allUpfront": "0.0677",
            "yrTerm1.allUpfront": "0.0781",
            "yrTerm1.partialUpfront": "0.0797",
            "yrTerm3.partialUpfront": "0.072"
          },
          "ondemand": "0.136"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.032",
            "yrTerm3.allUpfront": "0.0164",
            "yrTerm1.allUpfront": "0.0265",
            "yrTerm1.partialUpfront": "0.027",
            "yrTerm3.partialUpfront": "0.0174"
          },
          "ondemand": "0.036"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.018",
            "yrTerm3.allUpfront": "0.0115",
            "yrTerm1.allUpfront": "0.0172",
            "yrTerm1.partialUpfront": "0.0176",
            "yrTerm3.partialUpfront": "0.0123"
          },
          "ondemand": "0.026"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.105",
            "yrTerm3.allUpfront": "0.0753",
            "yrTerm1.allUpfront": "0.0877",
            "yrTerm1.partialUpfront": "0.0894",
            "yrTerm3.partialUpfront": "0.0801"
          },
          "ondemand": "0.15"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.043",
            "yrTerm3.allUpfront": "0.0236",
            "yrTerm1.allUpfront": "0.0362",
            "yrTerm1.partialUpfront": "0.0369",
            "yrTerm3.partialUpfront": "0.0251"
          },
          "ondemand": "0.05"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.03",
            "yrTerm3.allUpfront": "0.0188",
            "yrTerm1.allUpfront": "0.0288",
            "yrTerm1.partialUpfront": "0.0294",
            "yrTerm3.partialUpfront": "0.02"
          },
          "ondemand": "0.04"
        }
      },
      "sa-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.108",
            "yrTerm3.allUpfront": "0.0796",
            "yrTerm1.allUpfront": "0.0906",
            "yrTerm1.partialUpfront": "0.0925",
            "yrTerm3.partialUpfront": "0.0847"
          },
          "ondemand": "0.164"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.047",
            "yrTerm3.allUpfront": "0.0277",
            "yrTerm1.allUpfront": "0.0395",
            "yrTerm1.partialUpfront": "0.0403",
            "yrTerm3.partialUpfront": "0.0294"
          },
          "ondemand": "0.064"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.036",
            "yrTerm3.allUpfront": "0.0228",
            "yrTerm1.allUpfront": "0.0309",
            "yrTerm1.partialUpfront": "0.0316",
            "yrTerm3.partialUpfront": "0.0243"
          },
          "ondemand": "0.054"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.1444",
            "yrTerm3.allUpfront": "0.1288",
            "yrTerm1.allUpfront": "0.1388",
            "yrTerm1.partialUpfront": "0.1394",
            "yrTerm3.partialUpfront": "0.13"
          },
          "ondemand": "0.15"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.0436",
            "yrTerm3.allUpfront": "0.028",
            "yrTerm1.allUpfront": "0.038",
            "yrTerm1.partialUpfront": "0.0386",
            "yrTerm3.partialUpfront": "0.0292"
          },
          "ondemand": "0.049"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.0344",
            "yrTerm3.allUpfront": "0.0188",
            "yrTerm1.allUpfront": "0.0288",
            "yrTerm1.partialUpfront": "0.0294",
            "yrTerm3.partialUpfront": "0.02"
          },
          "ondemand": "0.04"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.107",
            "yrTerm3.allUpfront": "0.0787",
            "yrTerm1.allUpfront": "0.09",
            "yrTerm1.partialUpfront": "0.0918",
            "yrTerm3.partialUpfront": "0.0837"
          },
          "ondemand": "0.15"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.044",
            "yrTerm3.allUpfront": "0.0265",
            "yrTerm1.allUpfront": "0.0366",
            "yrTerm1.partialUpfront": "0.0374",
            "yrTerm3.partialUpfront": "0.0282"
          },
          "ondemand": "0.05"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.032",
            "yrTerm3.allUpfront": "0.0221",
            "yrTerm1.allUpfront": "0.0293",
            "yrTerm1.partialUpfront": "0.03",
            "yrTerm3.partialUpfront": "0.0235"
          },
          "ondemand": "0.04"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.107",
            "yrTerm3.allUpfront": "0.0787",
            "yrTerm1.allUpfront": "0.09",
            "yrTerm1.partialUpfront": "0.0918",
            "yrTerm3.partialUpfront": "0.0837"
          },
          "ondemand": "0.15"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.044",
            "yrTerm3.allUpfront": "0.0265",
            "yrTerm1.allUpfront": "0.0366",
            "yrTerm1.partialUpfront": "0.0374",
            "yrTerm3.partialUpfront": "0.0282"
          },
          "ondemand": "0.05"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.032",
            "yrTerm3.allUpfront": "0.0221",
            "yrTerm1.allUpfront": "0.0293",
            "yrTerm1.partialUpfront": "0.03",
            "yrTerm3.partialUpfront": "0.0235"
          },
          "ondemand": "0.04"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.093",
            "yrTerm3.allUpfront": "0.0677",
            "yrTerm1.allUpfront": "0.0781",
            "yrTerm1.partialUpfront": "0.0797",
            "yrTerm3.partialUpfront": "0.072"
          },
          "ondemand": "0.136"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.032",
            "yrTerm3.allUpfront": "0.0164",
            "yrTerm1.allUpfront": "0.0265",
            "yrTerm1.partialUpfront": "0.027",
            "yrTerm3.partialUpfront": "0.0174"
          },
          "ondemand": "0.036"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.018",
            "yrTerm3.allUpfront": "0.0115",
            "yrTerm1.allUpfront": "0.0172",
            "yrTerm1.partialUpfront": "0.0176",
            "yrTerm3.partialUpfront": "0.0123"
          },
          "ondemand": "0.026"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.098",
            "yrTerm3.allUpfront": "0.0701",
            "yrTerm1.allUpfront": "0.0822",
            "yrTerm1.partialUpfront": "0.0839",
            "yrTerm3.partialUpfront": "0.0746"
          },
          "ondemand": "0.141"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.036",
            "yrTerm3.allUpfront": "0.0187",
            "yrTerm1.allUpfront": "0.0304",
            "yrTerm1.partialUpfront": "0.031",
            "yrTerm3.partialUpfront": "0.0199"
          },
          "ondemand": "0.041"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.022",
            "yrTerm3.allUpfront": "0.0131",
            "yrTerm1.allUpfront": "0.0204",
            "yrTerm1.partialUpfront": "0.0209",
            "yrTerm3.partialUpfront": "0.0139"
          },
          "ondemand": "0.031"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.099",
            "yrTerm3.allUpfront": "0.0724",
            "yrTerm1.allUpfront": "0.0831",
            "yrTerm1.partialUpfront": "0.0848",
            "yrTerm3.partialUpfront": "0.077"
          },
          "ondemand": "0.144"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.038",
            "yrTerm3.allUpfront": "0.0209",
            "yrTerm1.allUpfront": "0.0321",
            "yrTerm1.partialUpfront": "0.0327",
            "yrTerm3.partialUpfront": "0.0222"
          },
          "ondemand": "0.044"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.024",
            "yrTerm3.allUpfront": "0.0159",
            "yrTerm1.allUpfront": "0.0229",
            "yrTerm1.partialUpfront": "0.0235",
            "yrTerm3.partialUpfront": "0.017"
          },
          "ondemand": "0.034"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.096",
            "yrTerm3.allUpfront": "0.07",
            "yrTerm1.allUpfront": "0.0806",
            "yrTerm1.partialUpfront": "0.0822",
            "yrTerm3.partialUpfront": "0.0745"
          },
          "ondemand": "0.14"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.034",
            "yrTerm3.allUpfront": "0.0186",
            "yrTerm1.allUpfront": "0.0288",
            "yrTerm1.partialUpfront": "0.0294",
            "yrTerm3.partialUpfront": "0.0198"
          },
          "ondemand": "0.04"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.022",
            "yrTerm3.allUpfront": "0.0136",
            "yrTerm1.allUpfront": "0.0201",
            "yrTerm1.partialUpfront": "0.0206",
            "yrTerm3.partialUpfront": "0.0145"
          },
          "ondemand": "0.03"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.094",
            "yrTerm3.allUpfront": "0.0693",
            "yrTerm1.allUpfront": "0.0789",
            "yrTerm1.partialUpfront": "0.0805",
            "yrTerm3.partialUpfront": "0.0738"
          },
          "ondemand": "0.138"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.033",
            "yrTerm3.allUpfront": "0.017",
            "yrTerm1.allUpfront": "0.0274",
            "yrTerm1.partialUpfront": "0.028",
            "yrTerm3.partialUpfront": "0.0181"
          },
          "ondemand": "0.038"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.02",
            "yrTerm3.allUpfront": "0.0126",
            "yrTerm1.allUpfront": "0.0186",
            "yrTerm1.partialUpfront": "0.019",
            "yrTerm3.partialUpfront": "0.0135"
          },
          "ondemand": "0.028"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 4,
      "max_enis": 2
    },
    "arch": [
      "x86_64",
      "i386"
    ],
    "vpc_only": true,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": false,
    "storage": null,
    "instance_type": "t2.small",
    "ECU": "variable",
    "memory": 2.0,
    "ebs_max_bandwidth": 0
  },
  {
    "family": "General purpose",
    "enhanced_networking": false,
    "vCPU": 2,
    "generation": "current",
    "ebs_iops": 0,
    "network_performance": "Low to Moderate",
    "ebs_throughput": 0,
    "pretty_name": "T2 Medium",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.187",
            "yrTerm3.allUpfront": "0.1361",
            "yrTerm1.allUpfront": "0.1563",
            "yrTerm1.partialUpfront": "0.1594",
            "yrTerm3.partialUpfront": "0.1448"
          },
          "ondemand": "0.272"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.062",
            "yrTerm3.allUpfront": "0.0327",
            "yrTerm1.allUpfront": "0.0519",
            "yrTerm1.partialUpfront": "0.053",
            "yrTerm3.partialUpfront": "0.0348"
          },
          "ondemand": "0.072"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.036",
            "yrTerm3.allUpfront": "0.0231",
            "yrTerm1.allUpfront": "0.0345",
            "yrTerm1.partialUpfront": "0.0353",
            "yrTerm3.partialUpfront": "0.0246"
          },
          "ondemand": "0.052"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.208",
            "yrTerm3.allUpfront": "0.1505",
            "yrTerm1.allUpfront": "0.1743",
            "yrTerm1.partialUpfront": "0.1778",
            "yrTerm3.partialUpfront": "0.1601"
          },
          "ondemand": "0.3"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.085",
            "yrTerm3.allUpfront": "0.0471",
            "yrTerm1.allUpfront": "0.0713",
            "yrTerm1.partialUpfront": "0.0728",
            "yrTerm3.partialUpfront": "0.0501"
          },
          "ondemand": "0.1"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.06",
            "yrTerm3.allUpfront": "0.0376",
            "yrTerm1.allUpfront": "0.0575",
            "yrTerm1.partialUpfront": "0.0588",
            "yrTerm3.partialUpfront": "0.04"
          },
          "ondemand": "0.08"
        }
      },
      "sa-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.216",
            "yrTerm3.allUpfront": "0.1587",
            "yrTerm1.allUpfront": "0.1814",
            "yrTerm1.partialUpfront": "0.185",
            "yrTerm3.partialUpfront": "0.1688"
          },
          "ondemand": "0.328"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.094",
            "yrTerm3.allUpfront": "0.0553",
            "yrTerm1.allUpfront": "0.0789",
            "yrTerm1.partialUpfront": "0.0805",
            "yrTerm3.partialUpfront": "0.0588"
          },
          "ondemand": "0.128"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.072",
            "yrTerm3.allUpfront": "0.0456",
            "yrTerm1.allUpfront": "0.0619",
            "yrTerm1.partialUpfront": "0.0631",
            "yrTerm3.partialUpfront": "0.0485"
          },
          "ondemand": "0.108"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.3988",
            "yrTerm3.allUpfront": "0.3676",
            "yrTerm1.allUpfront": "0.3876",
            "yrTerm1.partialUpfront": "0.3888",
            "yrTerm3.partialUpfront": "0.37"
          },
          "ondemand": "0.41"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.0872",
            "yrTerm3.allUpfront": "0.056",
            "yrTerm1.allUpfront": "0.076",
            "yrTerm1.partialUpfront": "0.0772",
            "yrTerm3.partialUpfront": "0.0584"
          },
          "ondemand": "0.098"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.0688",
            "yrTerm3.allUpfront": "0.0376",
            "yrTerm1.allUpfront": "0.0576",
            "yrTerm1.partialUpfront": "0.0588",
            "yrTerm3.partialUpfront": "0.04"
          },
          "ondemand": "0.08"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.215",
            "yrTerm3.allUpfront": "0.1576",
            "yrTerm1.allUpfront": "0.1798",
            "yrTerm1.partialUpfront": "0.1835",
            "yrTerm3.partialUpfront": "0.1677"
          },
          "ondemand": "0.3"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.088",
            "yrTerm3.allUpfront": "0.0539",
            "yrTerm1.allUpfront": "0.0735",
            "yrTerm1.partialUpfront": "0.075",
            "yrTerm3.partialUpfront": "0.0573"
          },
          "ondemand": "0.1"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.064",
            "yrTerm3.allUpfront": "0.0443",
            "yrTerm1.allUpfront": "0.0588",
            "yrTerm1.partialUpfront": "0.0601",
            "yrTerm3.partialUpfront": "0.0471"
          },
          "ondemand": "0.08"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.215",
            "yrTerm3.allUpfront": "0.1576",
            "yrTerm1.allUpfront": "0.1798",
            "yrTerm1.partialUpfront": "0.1835",
            "yrTerm3.partialUpfront": "0.1677"
          },
          "ondemand": "0.3"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.088",
            "yrTerm3.allUpfront": "0.0539",
            "yrTerm1.allUpfront": "0.0735",
            "yrTerm1.partialUpfront": "0.075",
            "yrTerm3.partialUpfront": "0.0573"
          },
          "ondemand": "0.1"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.064",
            "yrTerm3.allUpfront": "0.0443",
            "yrTerm1.allUpfront": "0.0588",
            "yrTerm1.partialUpfront": "0.0601",
            "yrTerm3.partialUpfront": "0.0471"
          },
          "ondemand": "0.08"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.187",
            "yrTerm3.allUpfront": "0.1361",
            "yrTerm1.allUpfront": "0.1563",
            "yrTerm1.partialUpfront": "0.1594",
            "yrTerm3.partialUpfront": "0.1448"
          },
          "ondemand": "0.272"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.062",
            "yrTerm3.allUpfront": "0.0327",
            "yrTerm1.allUpfront": "0.0519",
            "yrTerm1.partialUpfront": "0.053",
            "yrTerm3.partialUpfront": "0.0348"
          },
          "ondemand": "0.072"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.036",
            "yrTerm3.allUpfront": "0.0231",
            "yrTerm1.allUpfront": "0.0345",
            "yrTerm1.partialUpfront": "0.0353",
            "yrTerm3.partialUpfront": "0.0246"
          },
          "ondemand": "0.052"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.195",
            "yrTerm3.allUpfront": "0.1407",
            "yrTerm1.allUpfront": "0.1632",
            "yrTerm1.partialUpfront": "0.1667",
            "yrTerm3.partialUpfront": "0.1497"
          },
          "ondemand": "0.282"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.07",
            "yrTerm3.allUpfront": "0.0373",
            "yrTerm1.allUpfront": "0.0589",
            "yrTerm1.partialUpfront": "0.0602",
            "yrTerm3.partialUpfront": "0.0397"
          },
          "ondemand": "0.082"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.044",
            "yrTerm3.allUpfront": "0.0271",
            "yrTerm1.allUpfront": "0.041",
            "yrTerm1.partialUpfront": "0.0419",
            "yrTerm3.partialUpfront": "0.0289"
          },
          "ondemand": "0.062"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.198",
            "yrTerm3.allUpfront": "0.1451",
            "yrTerm1.allUpfront": "0.1662",
            "yrTerm1.partialUpfront": "0.1696",
            "yrTerm3.partialUpfront": "0.1544"
          },
          "ondemand": "0.288"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.077",
            "yrTerm3.allUpfront": "0.0417",
            "yrTerm1.allUpfront": "0.0642",
            "yrTerm1.partialUpfront": "0.0655",
            "yrTerm3.partialUpfront": "0.0444"
          },
          "ondemand": "0.088"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.048",
            "yrTerm3.allUpfront": "0.0319",
            "yrTerm1.allUpfront": "0.046",
            "yrTerm1.partialUpfront": "0.0471",
            "yrTerm3.partialUpfront": "0.034"
          },
          "ondemand": "0.068"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.193",
            "yrTerm3.allUpfront": "0.1405",
            "yrTerm1.allUpfront": "0.1615",
            "yrTerm1.partialUpfront": "0.1649",
            "yrTerm3.partialUpfront": "0.1495"
          },
          "ondemand": "0.28"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.068",
            "yrTerm3.allUpfront": "0.0371",
            "yrTerm1.allUpfront": "0.0573",
            "yrTerm1.partialUpfront": "0.0585",
            "yrTerm3.partialUpfront": "0.0394"
          },
          "ondemand": "0.08"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.044",
            "yrTerm3.allUpfront": "0.0273",
            "yrTerm1.allUpfront": "0.0402",
            "yrTerm1.partialUpfront": "0.0411",
            "yrTerm3.partialUpfront": "0.029"
          },
          "ondemand": "0.06"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.189",
            "yrTerm3.allUpfront": "0.1383",
            "yrTerm1.allUpfront": "0.1587",
            "yrTerm1.partialUpfront": "0.1619",
            "yrTerm3.partialUpfront": "0.1471"
          },
          "ondemand": "0.276"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.067",
            "yrTerm3.allUpfront": "0.0349",
            "yrTerm1.allUpfront": "0.0558",
            "yrTerm1.partialUpfront": "0.057",
            "yrTerm3.partialUpfront": "0.0371"
          },
          "ondemand": "0.076"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.04",
            "yrTerm3.allUpfront": "0.0253",
            "yrTerm1.allUpfront": "0.0372",
            "yrTerm1.partialUpfront": "0.038",
            "yrTerm3.partialUpfront": "0.0269"
          },
          "ondemand": "0.056"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 6,
      "max_enis": 3
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": true,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": false,
    "storage": null,
    "instance_type": "t2.medium",
    "ECU": "variable",
    "memory": 4.0,
    "ebs_max_bandwidth": 0
  },
  {
    "family": "General purpose",
    "enhanced_networking": false,
    "vCPU": 2,
    "generation": "current",
    "ebs_iops": 0,
    "network_performance": "Low to Moderate",
    "ebs_throughput": 0,
    "pretty_name": "T2 Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.294",
            "yrTerm3.allUpfront": "0.2158",
            "yrTerm1.allUpfront": "0.2459",
            "yrTerm1.partialUpfront": "0.2512",
            "yrTerm3.partialUpfront": "0.2296"
          },
          "ondemand": "0.434"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.106",
            "yrTerm3.allUpfront": "0.0654",
            "yrTerm1.allUpfront": "0.0893",
            "yrTerm1.partialUpfront": "0.0909",
            "yrTerm3.partialUpfront": "0.0696"
          },
          "ondemand": "0.134"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.072",
            "yrTerm3.allUpfront": "0.0462",
            "yrTerm1.allUpfront": "0.0689",
            "yrTerm1.partialUpfront": "0.0706",
            "yrTerm3.partialUpfront": "0.0492"
          },
          "ondemand": "0.104"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.348",
            "yrTerm3.allUpfront": "0.2445",
            "yrTerm1.allUpfront": "0.2918",
            "yrTerm1.partialUpfront": "0.2979",
            "yrTerm3.partialUpfront": "0.2601"
          },
          "ondemand": "0.49"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.161",
            "yrTerm3.allUpfront": "0.0942",
            "yrTerm1.allUpfront": "0.1349",
            "yrTerm1.partialUpfront": "0.1379",
            "yrTerm3.partialUpfront": "0.1002"
          },
          "ondemand": "0.19"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.12",
            "yrTerm3.allUpfront": "0.0753",
            "yrTerm1.allUpfront": "0.1151",
            "yrTerm1.partialUpfront": "0.1176",
            "yrTerm3.partialUpfront": "0.0801"
          },
          "ondemand": "0.16"
        }
      },
      "sa-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.358",
            "yrTerm3.allUpfront": "0.2608",
            "yrTerm1.allUpfront": "0.3006",
            "yrTerm1.partialUpfront": "0.3064",
            "yrTerm3.partialUpfront": "0.2776"
          },
          "ondemand": "0.546"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.171",
            "yrTerm3.allUpfront": "0.1106",
            "yrTerm1.allUpfront": "0.1438",
            "yrTerm1.partialUpfront": "0.1464",
            "yrTerm3.partialUpfront": "0.1176"
          },
          "ondemand": "0.246"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.144",
            "yrTerm3.allUpfront": "0.0912",
            "yrTerm1.allUpfront": "0.1237",
            "yrTerm1.partialUpfront": "0.1263",
            "yrTerm3.partialUpfront": "0.0971"
          },
          "ondemand": "0.216"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.6876",
            "yrTerm3.allUpfront": "0.6253",
            "yrTerm1.allUpfront": "0.6653",
            "yrTerm1.partialUpfront": "0.6676",
            "yrTerm3.partialUpfront": "0.6301"
          },
          "ondemand": "0.71"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.1652",
            "yrTerm3.allUpfront": "0.1029",
            "yrTerm1.allUpfront": "0.1429",
            "yrTerm1.partialUpfront": "0.1452",
            "yrTerm3.partialUpfront": "0.1077"
          },
          "ondemand": "0.188"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.1376",
            "yrTerm3.allUpfront": "0.0753",
            "yrTerm1.allUpfront": "0.1153",
            "yrTerm1.partialUpfront": "0.1176",
            "yrTerm3.partialUpfront": "0.0801"
          },
          "ondemand": "0.16"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.353",
            "yrTerm3.allUpfront": "0.2586",
            "yrTerm1.allUpfront": "0.295",
            "yrTerm1.partialUpfront": "0.3011",
            "yrTerm3.partialUpfront": "0.2753"
          },
          "ondemand": "0.49"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.165",
            "yrTerm3.allUpfront": "0.1078",
            "yrTerm1.allUpfront": "0.1379",
            "yrTerm1.partialUpfront": "0.1409",
            "yrTerm3.partialUpfront": "0.1146"
          },
          "ondemand": "0.19"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.128",
            "yrTerm3.allUpfront": "0.0885",
            "yrTerm1.allUpfront": "0.1176",
            "yrTerm1.partialUpfront": "0.1201",
            "yrTerm3.partialUpfront": "0.0942"
          },
          "ondemand": "0.16"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.353",
            "yrTerm3.allUpfront": "0.2586",
            "yrTerm1.allUpfront": "0.295",
            "yrTerm1.partialUpfront": "0.3011",
            "yrTerm3.partialUpfront": "0.2753"
          },
          "ondemand": "0.49"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.165",
            "yrTerm3.allUpfront": "0.1078",
            "yrTerm1.allUpfront": "0.1379",
            "yrTerm1.partialUpfront": "0.1409",
            "yrTerm3.partialUpfront": "0.1146"
          },
          "ondemand": "0.19"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.128",
            "yrTerm3.allUpfront": "0.0885",
            "yrTerm1.allUpfront": "0.1176",
            "yrTerm1.partialUpfront": "0.1201",
            "yrTerm3.partialUpfront": "0.0942"
          },
          "ondemand": "0.16"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.294",
            "yrTerm3.allUpfront": "0.2158",
            "yrTerm1.allUpfront": "0.2459",
            "yrTerm1.partialUpfront": "0.2512",
            "yrTerm3.partialUpfront": "0.2296"
          },
          "ondemand": "0.434"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.106",
            "yrTerm3.allUpfront": "0.0654",
            "yrTerm1.allUpfront": "0.0893",
            "yrTerm1.partialUpfront": "0.0909",
            "yrTerm3.partialUpfront": "0.0696"
          },
          "ondemand": "0.134"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.072",
            "yrTerm3.allUpfront": "0.0462",
            "yrTerm1.allUpfront": "0.0689",
            "yrTerm1.partialUpfront": "0.0706",
            "yrTerm3.partialUpfront": "0.0492"
          },
          "ondemand": "0.104"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.31",
            "yrTerm3.allUpfront": "0.2247",
            "yrTerm1.allUpfront": "0.2596",
            "yrTerm1.partialUpfront": "0.2654",
            "yrTerm3.partialUpfront": "0.2393"
          },
          "ondemand": "0.454"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.123",
            "yrTerm3.allUpfront": "0.0746",
            "yrTerm1.allUpfront": "0.1032",
            "yrTerm1.partialUpfront": "0.1054",
            "yrTerm3.partialUpfront": "0.0794"
          },
          "ondemand": "0.154"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.088",
            "yrTerm3.allUpfront": "0.0543",
            "yrTerm1.allUpfront": "0.082",
            "yrTerm1.partialUpfront": "0.0837",
            "yrTerm3.partialUpfront": "0.0578"
          },
          "ondemand": "0.124"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.321",
            "yrTerm3.allUpfront": "0.2337",
            "yrTerm1.allUpfront": "0.2693",
            "yrTerm1.partialUpfront": "0.2746",
            "yrTerm3.partialUpfront": "0.2487"
          },
          "ondemand": "0.466"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.135",
            "yrTerm3.allUpfront": "0.0835",
            "yrTerm1.allUpfront": "0.1127",
            "yrTerm1.partialUpfront": "0.1154",
            "yrTerm3.partialUpfront": "0.0888"
          },
          "ondemand": "0.166"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.096",
            "yrTerm3.allUpfront": "0.0638",
            "yrTerm1.allUpfront": "0.092",
            "yrTerm1.partialUpfront": "0.0941",
            "yrTerm3.partialUpfront": "0.0679"
          },
          "ondemand": "0.136"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.308",
            "yrTerm3.allUpfront": "0.2243",
            "yrTerm1.allUpfront": "0.2575",
            "yrTerm1.partialUpfront": "0.2626",
            "yrTerm3.partialUpfront": "0.2388"
          },
          "ondemand": "0.45"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.12",
            "yrTerm3.allUpfront": "0.074",
            "yrTerm1.allUpfront": "0.1009",
            "yrTerm1.partialUpfront": "0.1034",
            "yrTerm3.partialUpfront": "0.0788"
          },
          "ondemand": "0.15"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.088",
            "yrTerm3.allUpfront": "0.0546",
            "yrTerm1.allUpfront": "0.0804",
            "yrTerm1.partialUpfront": "0.0822",
            "yrTerm3.partialUpfront": "0.0581"
          },
          "ondemand": "0.12"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.299",
            "yrTerm3.allUpfront": "0.2201",
            "yrTerm1.allUpfront": "0.2511",
            "yrTerm1.partialUpfront": "0.2564",
            "yrTerm3.partialUpfront": "0.2342"
          },
          "ondemand": "0.442"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.113",
            "yrTerm3.allUpfront": "0.0698",
            "yrTerm1.allUpfront": "0.0944",
            "yrTerm1.partialUpfront": "0.0965",
            "yrTerm3.partialUpfront": "0.0743"
          },
          "ondemand": "0.142"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.08",
            "yrTerm3.allUpfront": "0.0505",
            "yrTerm1.allUpfront": "0.0744",
            "yrTerm1.partialUpfront": "0.0761",
            "yrTerm3.partialUpfront": "0.0538"
          },
          "ondemand": "0.112"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 12,
      "max_enis": 3
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": true,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": false,
    "storage": null,
    "instance_type": "t2.large",
    "ECU": "variable",
    "memory": 8.0,
    "ebs_max_bandwidth": 0
  },
  {
    "family": "General purpose",
    "enhanced_networking": true,
    "vCPU": 2,
    "generation": "current",
    "ebs_iops": 56.25,
    "network_performance": "Moderate",
    "ebs_throughput": 3600.0,
    "pretty_name": "M4 Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.194",
            "yrTerm3.allUpfront": "0.1335",
            "yrTerm1.allUpfront": "0.1627",
            "yrTerm1.partialUpfront": "0.166",
            "yrTerm3.partialUpfront": "0.142"
          },
          "ondemand": "0.24"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.815",
            "yrTerm3.allUpfront": "0.6331",
            "yrTerm1.allUpfront": "0.6823",
            "yrTerm1.partialUpfront": "0.6962",
            "yrTerm3.partialUpfront": "0.6738"
          },
          "ondemand": "0.921"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.184",
            "yrTerm3.allUpfront": "0.1219",
            "yrTerm1.allUpfront": "0.1579",
            "yrTerm1.partialUpfront": "0.1616",
            "yrTerm3.partialUpfront": "0.1298"
          },
          "ondemand": "0.246"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.083",
            "yrTerm3.allUpfront": "0.0446",
            "yrTerm1.allUpfront": "0.0688",
            "yrTerm1.partialUpfront": "0.0702",
            "yrTerm3.partialUpfront": "0.0477"
          },
          "ondemand": "0.12"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.222",
            "yrTerm3.allUpfront": "0.1532",
            "yrTerm1.allUpfront": "0.1862",
            "yrTerm1.partialUpfront": "0.19",
            "yrTerm3.partialUpfront": "0.1625"
          },
          "ondemand": "0.294"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.841",
            "yrTerm3.allUpfront": "0.6518",
            "yrTerm1.allUpfront": "0.7047",
            "yrTerm1.partialUpfront": "0.7196",
            "yrTerm3.partialUpfront": "0.6927"
          },
          "ondemand": "0.975"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.207",
            "yrTerm3.allUpfront": "0.1427",
            "yrTerm1.allUpfront": "0.1784",
            "yrTerm1.partialUpfront": "0.1821",
            "yrTerm3.partialUpfront": "0.1519"
          },
          "ondemand": "0.273"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.109",
            "yrTerm3.allUpfront": "0.0634",
            "yrTerm1.allUpfront": "0.0912",
            "yrTerm1.partialUpfront": "0.0936",
            "yrTerm3.partialUpfront": "0.0667"
          },
          "ondemand": "0.174"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.2724",
            "yrTerm3.allUpfront": "0.2288",
            "yrTerm1.allUpfront": "0.2556",
            "yrTerm1.partialUpfront": "0.2574",
            "yrTerm3.partialUpfront": "0.2327"
          },
          "ondemand": "0.334"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.9324",
            "yrTerm3.allUpfront": "0.8888",
            "yrTerm1.allUpfront": "0.9156",
            "yrTerm1.partialUpfront": "0.9174",
            "yrTerm3.partialUpfront": "0.8927"
          },
          "ondemand": "0.994"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.1954",
            "yrTerm3.allUpfront": "0.1518",
            "yrTerm1.allUpfront": "0.1786",
            "yrTerm1.partialUpfront": "0.1804",
            "yrTerm3.partialUpfront": "0.1557"
          },
          "ondemand": "0.257"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.1034",
            "yrTerm3.allUpfront": "0.0598",
            "yrTerm1.allUpfront": "0.0866",
            "yrTerm1.partialUpfront": "0.0884",
            "yrTerm3.partialUpfront": "0.0637"
          },
          "ondemand": "0.165"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.239",
            "yrTerm3.allUpfront": "0.1561",
            "yrTerm1.allUpfront": "0.1999",
            "yrTerm1.partialUpfront": "0.2041",
            "yrTerm3.partialUpfront": "0.166"
          },
          "ondemand": "0.298"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.856",
            "yrTerm3.allUpfront": "0.6545",
            "yrTerm1.allUpfront": "0.7178",
            "yrTerm1.partialUpfront": "0.7321",
            "yrTerm3.partialUpfront": "0.6961"
          },
          "ondemand": "0.979"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.216",
            "yrTerm3.allUpfront": "0.1401",
            "yrTerm1.allUpfront": "0.1868",
            "yrTerm1.partialUpfront": "0.1903",
            "yrTerm3.partialUpfront": "0.1495"
          },
          "ondemand": "0.304"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.124",
            "yrTerm3.allUpfront": "0.0661",
            "yrTerm1.allUpfront": "0.1042",
            "yrTerm1.partialUpfront": "0.1062",
            "yrTerm3.partialUpfront": "0.0701"
          },
          "ondemand": "0.178"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.233",
            "yrTerm3.allUpfront": "0.1561",
            "yrTerm1.allUpfront": "0.195",
            "yrTerm1.partialUpfront": "0.1985",
            "yrTerm3.partialUpfront": "0.166"
          },
          "ondemand": "0.288"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.852",
            "yrTerm3.allUpfront": "0.6545",
            "yrTerm1.allUpfront": "0.7131",
            "yrTerm1.partialUpfront": "0.7269",
            "yrTerm3.partialUpfront": "0.6961"
          },
          "ondemand": "0.969"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.211",
            "yrTerm3.allUpfront": "0.1401",
            "yrTerm1.allUpfront": "0.1818",
            "yrTerm1.partialUpfront": "0.1858",
            "yrTerm3.partialUpfront": "0.1495"
          },
          "ondemand": "0.294"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.119",
            "yrTerm3.allUpfront": "0.0661",
            "yrTerm1.allUpfront": "0.0997",
            "yrTerm1.partialUpfront": "0.1009",
            "yrTerm3.partialUpfront": "0.0701"
          },
          "ondemand": "0.168"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.194",
            "yrTerm3.allUpfront": "0.1335",
            "yrTerm1.allUpfront": "0.1627",
            "yrTerm1.partialUpfront": "0.166",
            "yrTerm3.partialUpfront": "0.142"
          },
          "ondemand": "0.24"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.815",
            "yrTerm3.allUpfront": "0.6331",
            "yrTerm1.allUpfront": "0.6823",
            "yrTerm1.partialUpfront": "0.6962",
            "yrTerm3.partialUpfront": "0.6738"
          },
          "ondemand": "0.921"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.184",
            "yrTerm3.allUpfront": "0.1219",
            "yrTerm1.allUpfront": "0.1579",
            "yrTerm1.partialUpfront": "0.1616",
            "yrTerm3.partialUpfront": "0.1298"
          },
          "ondemand": "0.246"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.083",
            "yrTerm3.allUpfront": "0.0446",
            "yrTerm1.allUpfront": "0.0688",
            "yrTerm1.partialUpfront": "0.0702",
            "yrTerm3.partialUpfront": "0.0477"
          },
          "ondemand": "0.12"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.227",
            "yrTerm3.allUpfront": "0.1504",
            "yrTerm1.allUpfront": "0.1901",
            "yrTerm1.partialUpfront": "0.194",
            "yrTerm3.partialUpfront": "0.16"
          },
          "ondemand": "0.26"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.846",
            "yrTerm3.allUpfront": "0.6492",
            "yrTerm1.allUpfront": "0.7084",
            "yrTerm1.partialUpfront": "0.7225",
            "yrTerm3.partialUpfront": "0.6903"
          },
          "ondemand": "0.941"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.222",
            "yrTerm3.allUpfront": "0.1426",
            "yrTerm1.allUpfront": "0.187",
            "yrTerm1.partialUpfront": "0.1891",
            "yrTerm3.partialUpfront": "0.1518"
          },
          "ondemand": "0.266"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.113",
            "yrTerm3.allUpfront": "0.0607",
            "yrTerm1.allUpfront": "0.095",
            "yrTerm1.partialUpfront": "0.0965",
            "yrTerm3.partialUpfront": "0.0643"
          },
          "ondemand": "0.14"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.222",
            "yrTerm3.allUpfront": "0.1523",
            "yrTerm1.allUpfront": "0.1862",
            "yrTerm1.partialUpfront": "0.19",
            "yrTerm3.partialUpfront": "0.162"
          },
          "ondemand": "0.263"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.841",
            "yrTerm3.allUpfront": "0.651",
            "yrTerm1.allUpfront": "0.7047",
            "yrTerm1.partialUpfront": "0.7196",
            "yrTerm3.partialUpfront": "0.6922"
          },
          "ondemand": "0.944"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.21",
            "yrTerm3.allUpfront": "0.1408",
            "yrTerm1.allUpfront": "0.1811",
            "yrTerm1.partialUpfront": "0.1844",
            "yrTerm3.partialUpfront": "0.1499"
          },
          "ondemand": "0.269"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.109",
            "yrTerm3.allUpfront": "0.0625",
            "yrTerm1.allUpfront": "0.0912",
            "yrTerm1.partialUpfront": "0.0936",
            "yrTerm3.partialUpfront": "0.0663"
          },
          "ondemand": "0.143"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.213",
            "yrTerm3.allUpfront": "0.1467",
            "yrTerm1.allUpfront": "0.1783",
            "yrTerm1.partialUpfront": "0.182",
            "yrTerm3.partialUpfront": "0.156"
          },
          "ondemand": "0.252"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.833",
            "yrTerm3.allUpfront": "0.6456",
            "yrTerm1.allUpfront": "0.6973",
            "yrTerm1.partialUpfront": "0.7117",
            "yrTerm3.partialUpfront": "0.6864"
          },
          "ondemand": "0.933"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.209",
            "yrTerm3.allUpfront": "0.1398",
            "yrTerm1.allUpfront": "0.1758",
            "yrTerm1.partialUpfront": "0.1793",
            "yrTerm3.partialUpfront": "0.1484"
          },
          "ondemand": "0.244"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.1",
            "yrTerm3.allUpfront": "0.0572",
            "yrTerm1.allUpfront": "0.0838",
            "yrTerm1.partialUpfront": "0.0857",
            "yrTerm3.partialUpfront": "0.0604"
          },
          "ondemand": "0.132"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 10,
      "max_enis": 2
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": true,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": true,
    "storage": null,
    "instance_type": "m4.large",
    "ECU": 6.5,
    "memory": 8.0,
    "ebs_max_bandwidth": 450.0
  },
  {
    "family": "General purpose",
    "enhanced_networking": true,
    "vCPU": 4,
    "generation": "current",
    "ebs_iops": 93.75,
    "network_performance": "High",
    "ebs_throughput": 6000.0,
    "pretty_name": "M4 Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.346",
            "yrTerm3.allUpfront": "0.2331",
            "yrTerm1.allUpfront": "0.2901",
            "yrTerm1.partialUpfront": "0.2959",
            "yrTerm3.partialUpfront": "0.248"
          },
          "ondemand": "0.439"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.961",
            "yrTerm3.allUpfront": "0.7295",
            "yrTerm1.allUpfront": "0.8051",
            "yrTerm1.partialUpfront": "0.8218",
            "yrTerm3.partialUpfront": "0.777"
          },
          "ondemand": "1.11"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.366",
            "yrTerm3.allUpfront": "0.2447",
            "yrTerm1.allUpfront": "0.3145",
            "yrTerm1.partialUpfront": "0.3205",
            "yrTerm3.partialUpfront": "0.2601"
          },
          "ondemand": "0.491"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.164",
            "yrTerm3.allUpfront": "0.0893",
            "yrTerm1.allUpfront": "0.1378",
            "yrTerm1.partialUpfront": "0.1403",
            "yrTerm3.partialUpfront": "0.0955"
          },
          "ondemand": "0.239"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.402",
            "yrTerm3.allUpfront": "0.2726",
            "yrTerm1.allUpfront": "0.3371",
            "yrTerm1.partialUpfront": "0.344",
            "yrTerm3.partialUpfront": "0.29"
          },
          "ondemand": "0.553"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.015",
            "yrTerm3.allUpfront": "0.767",
            "yrTerm1.allUpfront": "0.8499",
            "yrTerm1.partialUpfront": "0.8676",
            "yrTerm3.partialUpfront": "0.816"
          },
          "ondemand": "1.219"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.414",
            "yrTerm3.allUpfront": "0.2851",
            "yrTerm1.allUpfront": "0.3561",
            "yrTerm1.partialUpfront": "0.3636",
            "yrTerm3.partialUpfront": "0.3036"
          },
          "ondemand": "0.546"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.218",
            "yrTerm3.allUpfront": "0.1268",
            "yrTerm1.allUpfront": "0.1825",
            "yrTerm1.partialUpfront": "0.186",
            "yrTerm3.partialUpfront": "0.1345"
          },
          "ondemand": "0.348"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.4637",
            "yrTerm3.allUpfront": "0.3776",
            "yrTerm1.allUpfront": "0.4302",
            "yrTerm1.partialUpfront": "0.4337",
            "yrTerm3.partialUpfront": "0.3853"
          },
          "ondemand": "0.588"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.0987",
            "yrTerm3.allUpfront": "1.0126",
            "yrTerm1.allUpfront": "1.0652",
            "yrTerm1.partialUpfront": "1.0687",
            "yrTerm3.partialUpfront": "1.0203"
          },
          "ondemand": "1.223"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.3907",
            "yrTerm3.allUpfront": "0.3046",
            "yrTerm1.allUpfront": "0.3572",
            "yrTerm1.partialUpfront": "0.3607",
            "yrTerm3.partialUpfront": "0.3123"
          },
          "ondemand": "0.515"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.2067",
            "yrTerm3.allUpfront": "0.1206",
            "yrTerm1.allUpfront": "0.1732",
            "yrTerm1.partialUpfront": "0.1767",
            "yrTerm3.partialUpfront": "0.1283"
          },
          "ondemand": "0.331"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.435",
            "yrTerm3.allUpfront": "0.2782",
            "yrTerm1.allUpfront": "0.3646",
            "yrTerm1.partialUpfront": "0.372",
            "yrTerm3.partialUpfront": "0.296"
          },
          "ondemand": "0.561"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.046",
            "yrTerm3.allUpfront": "0.7723",
            "yrTerm1.allUpfront": "0.8759",
            "yrTerm1.partialUpfront": "0.8929",
            "yrTerm3.partialUpfront": "0.8218"
          },
          "ondemand": "1.226"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.433",
            "yrTerm3.allUpfront": "0.2802",
            "yrTerm1.allUpfront": "0.3734",
            "yrTerm1.partialUpfront": "0.3805",
            "yrTerm3.partialUpfront": "0.298"
          },
          "ondemand": "0.607"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.249",
            "yrTerm3.allUpfront": "0.1322",
            "yrTerm1.allUpfront": "0.2086",
            "yrTerm1.partialUpfront": "0.2124",
            "yrTerm3.partialUpfront": "0.1403"
          },
          "ondemand": "0.355"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.424",
            "yrTerm3.allUpfront": "0.2782",
            "yrTerm1.allUpfront": "0.3548",
            "yrTerm1.partialUpfront": "0.3621",
            "yrTerm3.partialUpfront": "0.296"
          },
          "ondemand": "0.541"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.035",
            "yrTerm3.allUpfront": "0.7723",
            "yrTerm1.allUpfront": "0.8666",
            "yrTerm1.partialUpfront": "0.8851",
            "yrTerm3.partialUpfront": "0.8218"
          },
          "ondemand": "1.207"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.422",
            "yrTerm3.allUpfront": "0.2802",
            "yrTerm1.allUpfront": "0.3636",
            "yrTerm1.partialUpfront": "0.3705",
            "yrTerm3.partialUpfront": "0.298"
          },
          "ondemand": "0.588"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.238",
            "yrTerm3.allUpfront": "0.1322",
            "yrTerm1.allUpfront": "0.1992",
            "yrTerm1.partialUpfront": "0.2036",
            "yrTerm3.partialUpfront": "0.1403"
          },
          "ondemand": "0.336"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.346",
            "yrTerm3.allUpfront": "0.2331",
            "yrTerm1.allUpfront": "0.2901",
            "yrTerm1.partialUpfront": "0.2959",
            "yrTerm3.partialUpfront": "0.248"
          },
          "ondemand": "0.439"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.961",
            "yrTerm3.allUpfront": "0.7295",
            "yrTerm1.allUpfront": "0.8051",
            "yrTerm1.partialUpfront": "0.8218",
            "yrTerm3.partialUpfront": "0.777"
          },
          "ondemand": "1.11"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.366",
            "yrTerm3.allUpfront": "0.2447",
            "yrTerm1.allUpfront": "0.3145",
            "yrTerm1.partialUpfront": "0.3205",
            "yrTerm3.partialUpfront": "0.2601"
          },
          "ondemand": "0.491"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.164",
            "yrTerm3.allUpfront": "0.0893",
            "yrTerm1.allUpfront": "0.1378",
            "yrTerm1.partialUpfront": "0.1403",
            "yrTerm3.partialUpfront": "0.0955"
          },
          "ondemand": "0.239"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.412",
            "yrTerm3.allUpfront": "0.267",
            "yrTerm1.allUpfront": "0.345",
            "yrTerm1.partialUpfront": "0.352",
            "yrTerm3.partialUpfront": "0.284"
          },
          "ondemand": "0.481"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.023",
            "yrTerm3.allUpfront": "0.7616",
            "yrTerm1.allUpfront": "0.8573",
            "yrTerm1.partialUpfront": "0.8753",
            "yrTerm3.partialUpfront": "0.8111"
          },
          "ondemand": "1.15"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.443",
            "yrTerm3.allUpfront": "0.286",
            "yrTerm1.allUpfront": "0.3739",
            "yrTerm1.partialUpfront": "0.3783",
            "yrTerm3.partialUpfront": "0.3042"
          },
          "ondemand": "0.531"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.227",
            "yrTerm3.allUpfront": "0.1215",
            "yrTerm1.allUpfront": "0.1898",
            "yrTerm1.partialUpfront": "0.1939",
            "yrTerm3.partialUpfront": "0.1296"
          },
          "ondemand": "0.279"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.402",
            "yrTerm3.allUpfront": "0.2707",
            "yrTerm1.allUpfront": "0.3371",
            "yrTerm1.partialUpfront": "0.344",
            "yrTerm3.partialUpfront": "0.288"
          },
          "ondemand": "0.487"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.015",
            "yrTerm3.allUpfront": "0.7651",
            "yrTerm1.allUpfront": "0.8499",
            "yrTerm1.partialUpfront": "0.8676",
            "yrTerm3.partialUpfront": "0.814"
          },
          "ondemand": "1.156"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.42",
            "yrTerm3.allUpfront": "0.2823",
            "yrTerm1.allUpfront": "0.3614",
            "yrTerm1.partialUpfront": "0.3684",
            "yrTerm3.partialUpfront": "0.3002"
          },
          "ondemand": "0.537"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.218",
            "yrTerm3.allUpfront": "0.125",
            "yrTerm1.allUpfront": "0.1825",
            "yrTerm1.partialUpfront": "0.186",
            "yrTerm3.partialUpfront": "0.1335"
          },
          "ondemand": "0.285"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.384",
            "yrTerm3.allUpfront": "0.2594",
            "yrTerm1.allUpfront": "0.3215",
            "yrTerm1.partialUpfront": "0.328",
            "yrTerm3.partialUpfront": "0.276"
          },
          "ondemand": "0.465"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.996",
            "yrTerm3.allUpfront": "0.7545",
            "yrTerm1.allUpfront": "0.835",
            "yrTerm1.partialUpfront": "0.853",
            "yrTerm3.partialUpfront": "0.8033"
          },
          "ondemand": "1.135"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.416",
            "yrTerm3.allUpfront": "0.2795",
            "yrTerm1.allUpfront": "0.3516",
            "yrTerm1.partialUpfront": "0.356",
            "yrTerm3.partialUpfront": "0.2977"
          },
          "ondemand": "0.488"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.2",
            "yrTerm3.allUpfront": "0.1143",
            "yrTerm1.allUpfront": "0.1676",
            "yrTerm1.partialUpfront": "0.1715",
            "yrTerm3.partialUpfront": "0.1218"
          },
          "ondemand": "0.264"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 15,
      "max_enis": 4
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": true,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": true,
    "storage": null,
    "instance_type": "m4.xlarge",
    "ECU": 13.0,
    "memory": 16.0,
    "ebs_max_bandwidth": 750.0
  },
  {
    "family": "General purpose",
    "enhanced_networking": true,
    "vCPU": 8,
    "generation": "current",
    "ebs_iops": 125.0,
    "network_performance": "High",
    "ebs_throughput": 8000.0,
    "pretty_name": "M4 Double Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.689",
            "yrTerm3.allUpfront": "0.4634",
            "yrTerm1.allUpfront": "0.5772",
            "yrTerm1.partialUpfront": "0.5885",
            "yrTerm3.partialUpfront": "0.4935"
          },
          "ondemand": "0.9"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.923",
            "yrTerm3.allUpfront": "1.4589",
            "yrTerm1.allUpfront": "1.6104",
            "yrTerm1.partialUpfront": "1.6436",
            "yrTerm3.partialUpfront": "1.552"
          },
          "ondemand": "2.321"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.735",
            "yrTerm3.allUpfront": "0.4895",
            "yrTerm1.allUpfront": "0.6298",
            "yrTerm1.partialUpfront": "0.6423",
            "yrTerm3.partialUpfront": "0.5204"
          },
          "ondemand": "0.983"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.329",
            "yrTerm3.allUpfront": "0.1786",
            "yrTerm1.allUpfront": "0.2756",
            "yrTerm1.partialUpfront": "0.2815",
            "yrTerm3.partialUpfront": "0.19"
          },
          "ondemand": "0.479"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.801",
            "yrTerm3.allUpfront": "0.5424",
            "yrTerm1.allUpfront": "0.6713",
            "yrTerm1.partialUpfront": "0.6855",
            "yrTerm3.partialUpfront": "0.5765"
          },
          "ondemand": "1.128"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.029",
            "yrTerm3.allUpfront": "1.5339",
            "yrTerm1.allUpfront": "1.6998",
            "yrTerm1.partialUpfront": "1.7342",
            "yrTerm3.partialUpfront": "1.6319"
          },
          "ondemand": "2.537"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.824",
            "yrTerm3.allUpfront": "0.5677",
            "yrTerm1.allUpfront": "0.7097",
            "yrTerm1.partialUpfront": "0.7241",
            "yrTerm3.partialUpfront": "0.6039"
          },
          "ondemand": "1.091"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.436",
            "yrTerm3.allUpfront": "0.2536",
            "yrTerm1.allUpfront": "0.365",
            "yrTerm1.partialUpfront": "0.3722",
            "yrTerm3.partialUpfront": "0.2699"
          },
          "ondemand": "0.695"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.9275",
            "yrTerm3.allUpfront": "0.7551",
            "yrTerm1.allUpfront": "0.8603",
            "yrTerm1.partialUpfront": "0.8674",
            "yrTerm3.partialUpfront": "0.7705"
          },
          "ondemand": "1.174"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.1975",
            "yrTerm3.allUpfront": "2.0251",
            "yrTerm1.allUpfront": "2.1303",
            "yrTerm1.partialUpfront": "2.1374",
            "yrTerm3.partialUpfront": "2.0405"
          },
          "ondemand": "2.444"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.7815",
            "yrTerm3.allUpfront": "0.6091",
            "yrTerm1.allUpfront": "0.7143",
            "yrTerm1.partialUpfront": "0.7214",
            "yrTerm3.partialUpfront": "0.6245"
          },
          "ondemand": "1.028"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.4135",
            "yrTerm3.allUpfront": "0.2411",
            "yrTerm1.allUpfront": "0.3463",
            "yrTerm1.partialUpfront": "0.3534",
            "yrTerm3.partialUpfront": "0.2565"
          },
          "ondemand": "0.66"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.867",
            "yrTerm3.allUpfront": "0.5537",
            "yrTerm1.allUpfront": "0.7261",
            "yrTerm1.partialUpfront": "0.7405",
            "yrTerm3.partialUpfront": "0.5895"
          },
          "ondemand": "1.144"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.092",
            "yrTerm3.allUpfront": "1.5446",
            "yrTerm1.allUpfront": "1.7518",
            "yrTerm1.partialUpfront": "1.7878",
            "yrTerm3.partialUpfront": "1.6436"
          },
          "ondemand": "2.553"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.867",
            "yrTerm3.allUpfront": "0.5597",
            "yrTerm1.allUpfront": "0.7477",
            "yrTerm1.partialUpfront": "0.7625",
            "yrTerm3.partialUpfront": "0.5957"
          },
          "ondemand": "1.215"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.498",
            "yrTerm3.allUpfront": "0.2643",
            "yrTerm1.allUpfront": "0.4171",
            "yrTerm1.partialUpfront": "0.4258",
            "yrTerm3.partialUpfront": "0.2816"
          },
          "ondemand": "0.711"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.844",
            "yrTerm3.allUpfront": "0.5537",
            "yrTerm1.allUpfront": "0.7066",
            "yrTerm1.partialUpfront": "0.7205",
            "yrTerm3.partialUpfront": "0.5895"
          },
          "ondemand": "1.104"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.069",
            "yrTerm3.allUpfront": "1.5446",
            "yrTerm1.allUpfront": "1.7332",
            "yrTerm1.partialUpfront": "1.7682",
            "yrTerm3.partialUpfront": "1.6436"
          },
          "ondemand": "2.515"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.844",
            "yrTerm3.allUpfront": "0.5597",
            "yrTerm1.allUpfront": "0.7282",
            "yrTerm1.partialUpfront": "0.7435",
            "yrTerm3.partialUpfront": "0.5957"
          },
          "ondemand": "1.177"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.476",
            "yrTerm3.allUpfront": "0.2643",
            "yrTerm1.allUpfront": "0.3984",
            "yrTerm1.partialUpfront": "0.4063",
            "yrTerm3.partialUpfront": "0.2816"
          },
          "ondemand": "0.673"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.689",
            "yrTerm3.allUpfront": "0.4634",
            "yrTerm1.allUpfront": "0.5772",
            "yrTerm1.partialUpfront": "0.5885",
            "yrTerm3.partialUpfront": "0.4935"
          },
          "ondemand": "0.9"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.923",
            "yrTerm3.allUpfront": "1.4589",
            "yrTerm1.allUpfront": "1.6104",
            "yrTerm1.partialUpfront": "1.6436",
            "yrTerm3.partialUpfront": "1.552"
          },
          "ondemand": "2.321"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.735",
            "yrTerm3.allUpfront": "0.4895",
            "yrTerm1.allUpfront": "0.6298",
            "yrTerm1.partialUpfront": "0.6423",
            "yrTerm3.partialUpfront": "0.5204"
          },
          "ondemand": "0.983"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.329",
            "yrTerm3.allUpfront": "0.1786",
            "yrTerm1.allUpfront": "0.2756",
            "yrTerm1.partialUpfront": "0.2815",
            "yrTerm3.partialUpfront": "0.19"
          },
          "ondemand": "0.479"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.82",
            "yrTerm3.allUpfront": "0.5311",
            "yrTerm1.allUpfront": "0.687",
            "yrTerm1.partialUpfront": "0.7015",
            "yrTerm3.partialUpfront": "0.5655"
          },
          "ondemand": "0.984"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.047",
            "yrTerm3.allUpfront": "1.5232",
            "yrTerm1.allUpfront": "1.7146",
            "yrTerm1.partialUpfront": "1.7499",
            "yrTerm3.partialUpfront": "1.6202"
          },
          "ondemand": "2.401"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.885",
            "yrTerm3.allUpfront": "0.5703",
            "yrTerm1.allUpfront": "0.7479",
            "yrTerm1.partialUpfront": "0.7562",
            "yrTerm3.partialUpfront": "0.6063"
          },
          "ondemand": "1.063"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.453",
            "yrTerm3.allUpfront": "0.2429",
            "yrTerm1.allUpfront": "0.3799",
            "yrTerm1.partialUpfront": "0.3878",
            "yrTerm3.partialUpfront": "0.2582"
          },
          "ondemand": "0.559"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.801",
            "yrTerm3.allUpfront": "0.5386",
            "yrTerm1.allUpfront": "0.6713",
            "yrTerm1.partialUpfront": "0.6855",
            "yrTerm3.partialUpfront": "0.5735"
          },
          "ondemand": "0.996"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.029",
            "yrTerm3.allUpfront": "1.5303",
            "yrTerm1.allUpfront": "1.6998",
            "yrTerm1.partialUpfront": "1.7342",
            "yrTerm3.partialUpfront": "1.628"
          },
          "ondemand": "2.412"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.842",
            "yrTerm3.allUpfront": "0.5646",
            "yrTerm1.allUpfront": "0.7243",
            "yrTerm1.partialUpfront": "0.7395",
            "yrTerm3.partialUpfront": "0.6003"
          },
          "ondemand": "1.074"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.436",
            "yrTerm3.allUpfront": "0.25",
            "yrTerm1.allUpfront": "0.365",
            "yrTerm1.partialUpfront": "0.3722",
            "yrTerm3.partialUpfront": "0.266"
          },
          "ondemand": "0.57"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.764",
            "yrTerm3.allUpfront": "0.5161",
            "yrTerm1.allUpfront": "0.64",
            "yrTerm1.partialUpfront": "0.6535",
            "yrTerm3.partialUpfront": "0.5485"
          },
          "ondemand": "0.952"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.994",
            "yrTerm3.allUpfront": "1.5089",
            "yrTerm1.allUpfront": "1.6699",
            "yrTerm1.partialUpfront": "1.7039",
            "yrTerm3.partialUpfront": "1.6056"
          },
          "ondemand": "2.37"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.835",
            "yrTerm3.allUpfront": "0.5572",
            "yrTerm1.allUpfront": "0.7032",
            "yrTerm1.partialUpfront": "0.7137",
            "yrTerm3.partialUpfront": "0.5923"
          },
          "ondemand": "0.976"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.4",
            "yrTerm3.allUpfront": "0.2286",
            "yrTerm1.allUpfront": "0.3352",
            "yrTerm1.partialUpfront": "0.342",
            "yrTerm3.partialUpfront": "0.2436"
          },
          "ondemand": "0.528"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 15,
      "max_enis": 4
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": true,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": true,
    "storage": null,
    "instance_type": "m4.2xlarge",
    "ECU": 26.0,
    "memory": 32.0,
    "ebs_max_bandwidth": 1000.0
  },
  {
    "family": "General purpose",
    "enhanced_networking": true,
    "vCPU": 16,
    "generation": "current",
    "ebs_iops": 250.0,
    "network_performance": "High",
    "ebs_throughput": 16000.0,
    "pretty_name": "M4 Quadruple Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.381",
            "yrTerm3.allUpfront": "0.9287",
            "yrTerm1.allUpfront": "1.1564",
            "yrTerm1.partialUpfront": "1.18",
            "yrTerm3.partialUpfront": "0.988"
          },
          "ondemand": "1.844"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.845",
            "yrTerm3.allUpfront": "2.9178",
            "yrTerm1.allUpfront": "3.2207",
            "yrTerm1.partialUpfront": "3.2862",
            "yrTerm3.partialUpfront": "3.104"
          },
          "ondemand": "4.594"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.469",
            "yrTerm3.allUpfront": "0.9789",
            "yrTerm1.allUpfront": "1.2595",
            "yrTerm1.partialUpfront": "1.2856",
            "yrTerm3.partialUpfront": "1.0417"
          },
          "ondemand": "1.966"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.658",
            "yrTerm3.allUpfront": "0.3572",
            "yrTerm1.allUpfront": "0.5511",
            "yrTerm1.partialUpfront": "0.5622",
            "yrTerm3.partialUpfront": "0.38"
          },
          "ondemand": "0.958"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.605",
            "yrTerm3.allUpfront": "1.0866",
            "yrTerm1.allUpfront": "1.3445",
            "yrTerm1.partialUpfront": "1.372",
            "yrTerm3.partialUpfront": "1.156"
          },
          "ondemand": "2.3"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "4.058",
            "yrTerm3.allUpfront": "3.0678",
            "yrTerm1.allUpfront": "3.3994",
            "yrTerm1.partialUpfront": "3.4684",
            "yrTerm3.partialUpfront": "3.2638"
          },
          "ondemand": "5.027"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.649",
            "yrTerm3.allUpfront": "1.1354",
            "yrTerm1.allUpfront": "1.4194",
            "yrTerm1.partialUpfront": "1.4482",
            "yrTerm3.partialUpfront": "1.2079"
          },
          "ondemand": "2.183"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.871",
            "yrTerm3.allUpfront": "0.5072",
            "yrTerm1.allUpfront": "0.7299",
            "yrTerm1.partialUpfront": "0.7444",
            "yrTerm3.partialUpfront": "0.5398"
          },
          "ondemand": "1.391"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.8561",
            "yrTerm3.allUpfront": "1.5102",
            "yrTerm1.allUpfront": "1.7216",
            "yrTerm1.partialUpfront": "1.7358",
            "yrTerm3.partialUpfront": "1.541"
          },
          "ondemand": "2.349"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "4.3961",
            "yrTerm3.allUpfront": "4.0502",
            "yrTerm1.allUpfront": "4.2616",
            "yrTerm1.partialUpfront": "4.2758",
            "yrTerm3.partialUpfront": "4.081"
          },
          "ondemand": "4.889"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.5641",
            "yrTerm3.allUpfront": "1.2182",
            "yrTerm1.allUpfront": "1.4296",
            "yrTerm1.partialUpfront": "1.4438",
            "yrTerm3.partialUpfront": "1.249"
          },
          "ondemand": "2.057"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.8281",
            "yrTerm3.allUpfront": "0.4822",
            "yrTerm1.allUpfront": "0.6936",
            "yrTerm1.partialUpfront": "0.7078",
            "yrTerm3.partialUpfront": "0.513"
          },
          "ondemand": "1.321"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.736",
            "yrTerm3.allUpfront": "1.1092",
            "yrTerm1.allUpfront": "1.4543",
            "yrTerm1.partialUpfront": "1.484",
            "yrTerm3.partialUpfront": "1.18"
          },
          "ondemand": "2.332"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "4.183",
            "yrTerm3.allUpfront": "3.0892",
            "yrTerm1.allUpfront": "3.5037",
            "yrTerm1.partialUpfront": "3.5757",
            "yrTerm3.partialUpfront": "3.2862"
          },
          "ondemand": "5.057"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.733",
            "yrTerm3.allUpfront": "1.1194",
            "yrTerm1.allUpfront": "1.4954",
            "yrTerm1.partialUpfront": "1.526",
            "yrTerm3.partialUpfront": "1.1904"
          },
          "ondemand": "2.429"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.996",
            "yrTerm3.allUpfront": "0.5287",
            "yrTerm1.allUpfront": "0.8341",
            "yrTerm1.partialUpfront": "0.8516",
            "yrTerm3.partialUpfront": "0.5622"
          },
          "ondemand": "1.421"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.689",
            "yrTerm3.allUpfront": "1.1092",
            "yrTerm1.allUpfront": "1.4151",
            "yrTerm1.partialUpfront": "1.444",
            "yrTerm3.partialUpfront": "1.18"
          },
          "ondemand": "2.252"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "4.139",
            "yrTerm3.allUpfront": "3.0892",
            "yrTerm1.allUpfront": "3.4666",
            "yrTerm1.partialUpfront": "3.5376",
            "yrTerm3.partialUpfront": "3.2862"
          },
          "ondemand": "4.981"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.689",
            "yrTerm3.allUpfront": "1.1194",
            "yrTerm1.allUpfront": "1.4563",
            "yrTerm1.partialUpfront": "1.486",
            "yrTerm3.partialUpfront": "1.1904"
          },
          "ondemand": "2.353"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.952",
            "yrTerm3.allUpfront": "0.5287",
            "yrTerm1.allUpfront": "0.797",
            "yrTerm1.partialUpfront": "0.8136",
            "yrTerm3.partialUpfront": "0.5622"
          },
          "ondemand": "1.345"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.381",
            "yrTerm3.allUpfront": "0.9287",
            "yrTerm1.allUpfront": "1.1564",
            "yrTerm1.partialUpfront": "1.18",
            "yrTerm3.partialUpfront": "0.988"
          },
          "ondemand": "1.844"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.845",
            "yrTerm3.allUpfront": "2.9178",
            "yrTerm1.allUpfront": "3.2207",
            "yrTerm1.partialUpfront": "3.2862",
            "yrTerm3.partialUpfront": "3.104"
          },
          "ondemand": "4.594"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.469",
            "yrTerm3.allUpfront": "0.9789",
            "yrTerm1.allUpfront": "1.2595",
            "yrTerm1.partialUpfront": "1.2856",
            "yrTerm3.partialUpfront": "1.0417"
          },
          "ondemand": "1.966"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.658",
            "yrTerm3.allUpfront": "0.3572",
            "yrTerm1.allUpfront": "0.5511",
            "yrTerm1.partialUpfront": "0.5622",
            "yrTerm3.partialUpfront": "0.38"
          },
          "ondemand": "0.958"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.643",
            "yrTerm3.allUpfront": "1.0641",
            "yrTerm1.allUpfront": "1.3759",
            "yrTerm1.partialUpfront": "1.4041",
            "yrTerm3.partialUpfront": "1.132"
          },
          "ondemand": "2.012"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "4.094",
            "yrTerm3.allUpfront": "3.0463",
            "yrTerm1.allUpfront": "3.4292",
            "yrTerm1.partialUpfront": "3.4995",
            "yrTerm3.partialUpfront": "3.2404"
          },
          "ondemand": "4.753"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.769",
            "yrTerm3.allUpfront": "1.1406",
            "yrTerm1.allUpfront": "1.4957",
            "yrTerm1.partialUpfront": "1.5132",
            "yrTerm3.partialUpfront": "1.2137"
          },
          "ondemand": "2.125"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.907",
            "yrTerm3.allUpfront": "0.4858",
            "yrTerm1.allUpfront": "0.7597",
            "yrTerm1.partialUpfront": "0.7756",
            "yrTerm3.partialUpfront": "0.5164"
          },
          "ondemand": "1.117"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.605",
            "yrTerm3.allUpfront": "1.0791",
            "yrTerm1.allUpfront": "1.3445",
            "yrTerm1.partialUpfront": "1.372",
            "yrTerm3.partialUpfront": "1.148"
          },
          "ondemand": "2.036"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "4.058",
            "yrTerm3.allUpfront": "3.0606",
            "yrTerm1.allUpfront": "3.3994",
            "yrTerm1.partialUpfront": "3.4684",
            "yrTerm3.partialUpfront": "3.256"
          },
          "ondemand": "4.776"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.683",
            "yrTerm3.allUpfront": "1.1293",
            "yrTerm1.allUpfront": "1.4486",
            "yrTerm1.partialUpfront": "1.4782",
            "yrTerm3.partialUpfront": "1.2017"
          },
          "ondemand": "2.148"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.871",
            "yrTerm3.allUpfront": "0.5001",
            "yrTerm1.allUpfront": "0.7299",
            "yrTerm1.partialUpfront": "0.7444",
            "yrTerm3.partialUpfront": "0.532"
          },
          "ondemand": "1.14"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.53",
            "yrTerm3.allUpfront": "1.034",
            "yrTerm1.allUpfront": "1.2818",
            "yrTerm1.partialUpfront": "1.308",
            "yrTerm3.partialUpfront": "1.1"
          },
          "ondemand": "1.948"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.987",
            "yrTerm3.allUpfront": "3.0178",
            "yrTerm1.allUpfront": "3.3398",
            "yrTerm1.partialUpfront": "3.408",
            "yrTerm3.partialUpfront": "3.2102"
          },
          "ondemand": "4.692"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.67",
            "yrTerm3.allUpfront": "1.1143",
            "yrTerm1.allUpfront": "1.4063",
            "yrTerm1.partialUpfront": "1.4276",
            "yrTerm3.partialUpfront": "1.1857"
          },
          "ondemand": "1.952"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.8",
            "yrTerm3.allUpfront": "0.4572",
            "yrTerm1.allUpfront": "0.6703",
            "yrTerm1.partialUpfront": "0.684",
            "yrTerm3.partialUpfront": "0.4862"
          },
          "ondemand": "1.056"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 30,
      "max_enis": 8
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": true,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": true,
    "storage": null,
    "instance_type": "m4.4xlarge",
    "ECU": 53.5,
    "memory": 64.0,
    "ebs_max_bandwidth": 2000.0
  },
  {
    "family": "General purpose",
    "enhanced_networking": true,
    "vCPU": 40,
    "generation": "current",
    "ebs_iops": 500.0,
    "network_performance": "10 Gigabit",
    "ebs_throughput": 32000.0,
    "pretty_name": "M4 Deca Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.449",
            "yrTerm3.allUpfront": "2.3199",
            "yrTerm1.allUpfront": "2.889",
            "yrTerm1.partialUpfront": "2.948",
            "yrTerm3.partialUpfront": "2.468"
          },
          "ondemand": "4.582"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "9.61",
            "yrTerm3.allUpfront": "7.2925",
            "yrTerm1.allUpfront": "8.0498",
            "yrTerm1.partialUpfront": "8.214",
            "yrTerm3.partialUpfront": "7.758"
          },
          "ondemand": "11.679"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "3.672",
            "yrTerm3.allUpfront": "2.4473",
            "yrTerm1.allUpfront": "3.1487",
            "yrTerm1.partialUpfront": "3.2125",
            "yrTerm3.partialUpfront": "2.6038"
          },
          "ondemand": "4.914"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.645",
            "yrTerm3.allUpfront": "0.893",
            "yrTerm1.allUpfront": "1.378",
            "yrTerm1.partialUpfront": "1.406",
            "yrTerm3.partialUpfront": "0.95"
          },
          "ondemand": "2.394"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "4.011",
            "yrTerm3.allUpfront": "2.7147",
            "yrTerm1.allUpfront": "3.3595",
            "yrTerm1.partialUpfront": "3.428",
            "yrTerm3.partialUpfront": "2.888"
          },
          "ondemand": "5.722"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "10.144",
            "yrTerm3.allUpfront": "7.6676",
            "yrTerm1.allUpfront": "8.4967",
            "yrTerm1.partialUpfront": "8.67",
            "yrTerm3.partialUpfront": "8.1575"
          },
          "ondemand": "12.762"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "4.122",
            "yrTerm3.allUpfront": "2.8383",
            "yrTerm1.allUpfront": "3.5486",
            "yrTerm1.partialUpfront": "3.6205",
            "yrTerm3.partialUpfront": "3.0197"
          },
          "ondemand": "5.457"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.178",
            "yrTerm3.allUpfront": "1.2681",
            "yrTerm1.allUpfront": "1.8248",
            "yrTerm1.partialUpfront": "1.8621",
            "yrTerm3.partialUpfront": "1.3495"
          },
          "ondemand": "3.477"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "4.6396",
            "yrTerm3.allUpfront": "3.7747",
            "yrTerm1.allUpfront": "4.3035",
            "yrTerm1.partialUpfront": "4.3389",
            "yrTerm3.partialUpfront": "3.8516"
          },
          "ondemand": "5.873"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "10.9896",
            "yrTerm3.allUpfront": "10.1247",
            "yrTerm1.allUpfront": "10.6535",
            "yrTerm1.partialUpfront": "10.6889",
            "yrTerm3.partialUpfront": "10.2016"
          },
          "ondemand": "12.223"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "3.9096",
            "yrTerm3.allUpfront": "3.0447",
            "yrTerm1.allUpfront": "3.5735",
            "yrTerm1.partialUpfront": "3.6089",
            "yrTerm3.partialUpfront": "3.1216"
          },
          "ondemand": "5.143"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.0696",
            "yrTerm3.allUpfront": "1.2047",
            "yrTerm1.allUpfront": "1.7335",
            "yrTerm1.partialUpfront": "1.7689",
            "yrTerm3.partialUpfront": "1.2816"
          },
          "ondemand": "3.303"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "4.338",
            "yrTerm3.allUpfront": "2.7711",
            "yrTerm1.allUpfront": "3.6338",
            "yrTerm1.partialUpfront": "3.708",
            "yrTerm3.partialUpfront": "2.948"
          },
          "ondemand": "5.802"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "10.455",
            "yrTerm3.allUpfront": "7.7212",
            "yrTerm1.allUpfront": "8.7573",
            "yrTerm1.partialUpfront": "8.9359",
            "yrTerm3.partialUpfront": "8.214"
          },
          "ondemand": "12.838"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "4.333",
            "yrTerm3.allUpfront": "2.7984",
            "yrTerm1.allUpfront": "3.7387",
            "yrTerm1.partialUpfront": "3.8145",
            "yrTerm3.partialUpfront": "2.9765"
          },
          "ondemand": "6.073"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.49",
            "yrTerm3.allUpfront": "1.3217",
            "yrTerm1.allUpfront": "2.0855",
            "yrTerm1.partialUpfront": "2.1279",
            "yrTerm3.partialUpfront": "1.406"
          },
          "ondemand": "3.553"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "4.221",
            "yrTerm3.allUpfront": "2.7711",
            "yrTerm1.allUpfront": "3.5358",
            "yrTerm1.partialUpfront": "3.608",
            "yrTerm3.partialUpfront": "2.948"
          },
          "ondemand": "5.602"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "10.344",
            "yrTerm3.allUpfront": "7.7212",
            "yrTerm1.allUpfront": "8.6642",
            "yrTerm1.partialUpfront": "8.8414",
            "yrTerm3.partialUpfront": "8.214"
          },
          "ondemand": "12.648"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "4.222",
            "yrTerm3.allUpfront": "2.7984",
            "yrTerm1.allUpfront": "3.6408",
            "yrTerm1.partialUpfront": "3.7145",
            "yrTerm3.partialUpfront": "2.9765"
          },
          "ondemand": "5.883"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.379",
            "yrTerm3.allUpfront": "1.3217",
            "yrTerm1.allUpfront": "1.9922",
            "yrTerm1.partialUpfront": "2.0334",
            "yrTerm3.partialUpfront": "1.406"
          },
          "ondemand": "3.363"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.449",
            "yrTerm3.allUpfront": "2.3199",
            "yrTerm1.allUpfront": "2.889",
            "yrTerm1.partialUpfront": "2.948",
            "yrTerm3.partialUpfront": "2.468"
          },
          "ondemand": "4.582"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "9.61",
            "yrTerm3.allUpfront": "7.2925",
            "yrTerm1.allUpfront": "8.0498",
            "yrTerm1.partialUpfront": "8.214",
            "yrTerm3.partialUpfront": "7.758"
          },
          "ondemand": "11.679"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "3.672",
            "yrTerm3.allUpfront": "2.4473",
            "yrTerm1.allUpfront": "3.1487",
            "yrTerm1.partialUpfront": "3.2125",
            "yrTerm3.partialUpfront": "2.6038"
          },
          "ondemand": "4.914"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.645",
            "yrTerm3.allUpfront": "0.893",
            "yrTerm1.allUpfront": "1.378",
            "yrTerm1.partialUpfront": "1.406",
            "yrTerm3.partialUpfront": "0.95"
          },
          "ondemand": "2.394"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "4.104",
            "yrTerm3.allUpfront": "2.6583",
            "yrTerm1.allUpfront": "3.4378",
            "yrTerm1.partialUpfront": "3.508",
            "yrTerm3.partialUpfront": "2.828"
          },
          "ondemand": "5.002"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "10.233",
            "yrTerm3.allUpfront": "7.614",
            "yrTerm1.allUpfront": "8.571",
            "yrTerm1.partialUpfront": "8.7459",
            "yrTerm3.partialUpfront": "8.1"
          },
          "ondemand": "12.078"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "4.424",
            "yrTerm3.allUpfront": "2.8515",
            "yrTerm1.allUpfront": "3.7393",
            "yrTerm1.partialUpfront": "3.7805",
            "yrTerm3.partialUpfront": "3.0337"
          },
          "ondemand": "5.313"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.268",
            "yrTerm3.allUpfront": "1.2145",
            "yrTerm1.allUpfront": "1.8992",
            "yrTerm1.partialUpfront": "1.9379",
            "yrTerm3.partialUpfront": "1.292"
          },
          "ondemand": "2.793"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "4.011",
            "yrTerm3.allUpfront": "2.6959",
            "yrTerm1.allUpfront": "3.3595",
            "yrTerm1.partialUpfront": "3.428",
            "yrTerm3.partialUpfront": "2.868"
          },
          "ondemand": "5.062"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "10.144",
            "yrTerm3.allUpfront": "7.6497",
            "yrTerm1.allUpfront": "8.4967",
            "yrTerm1.partialUpfront": "8.67",
            "yrTerm3.partialUpfront": "8.138"
          },
          "ondemand": "12.135"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "4.209",
            "yrTerm3.allUpfront": "2.8233",
            "yrTerm1.allUpfront": "3.6216",
            "yrTerm1.partialUpfront": "3.6957",
            "yrTerm3.partialUpfront": "3.0038"
          },
          "ondemand": "5.37"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.178",
            "yrTerm3.allUpfront": "1.2502",
            "yrTerm1.allUpfront": "1.8248",
            "yrTerm1.partialUpfront": "1.8621",
            "yrTerm3.partialUpfront": "1.33"
          },
          "ondemand": "2.85"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.824",
            "yrTerm3.allUpfront": "2.5831",
            "yrTerm1.allUpfront": "3.2026",
            "yrTerm1.partialUpfront": "3.268",
            "yrTerm3.partialUpfront": "2.748"
          },
          "ondemand": "4.842"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "9.966",
            "yrTerm3.allUpfront": "7.5425",
            "yrTerm1.allUpfront": "8.3476",
            "yrTerm1.partialUpfront": "8.518",
            "yrTerm3.partialUpfront": "8.024"
          },
          "ondemand": "11.926"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "4.175",
            "yrTerm3.allUpfront": "2.7857",
            "yrTerm1.allUpfront": "3.5159",
            "yrTerm1.partialUpfront": "3.568",
            "yrTerm3.partialUpfront": "2.9637"
          },
          "ondemand": "4.881"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.001",
            "yrTerm3.allUpfront": "1.143",
            "yrTerm1.allUpfront": "1.6758",
            "yrTerm1.partialUpfront": "1.71",
            "yrTerm3.partialUpfront": "1.216"
          },
          "ondemand": "2.641"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 30,
      "max_enis": 8
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": true,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": true,
    "storage": null,
    "instance_type": "m4.10xlarge",
    "ECU": 124.5,
    "memory": 160.0,
    "ebs_max_bandwidth": 4000.0
  },
  {
    "family": "General purpose",
    "enhanced_networking": false,
    "vCPU": 1,
    "generation": "current",
    "ebs_iops": 0,
    "network_performance": "Moderate",
    "ebs_throughput": 0,
    "pretty_name": "M3 General Purpose Medium",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.127",
            "yrTerm3.allUpfront": "0.0785",
            "yrTerm1.allUpfront": "0.1059",
            "yrTerm1.partialUpfront": "0.1082",
            "yrTerm3.partialUpfront": "0.0836"
          },
          "ondemand": "0.184"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.221",
            "yrTerm3.allUpfront": "0.155",
            "yrTerm1.allUpfront": "0.184",
            "yrTerm1.partialUpfront": "0.1883",
            "yrTerm3.partialUpfront": "0.1649"
          },
          "ondemand": "0.353"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.1",
            "yrTerm3.allUpfront": "0.0641",
            "yrTerm1.allUpfront": "0.084",
            "yrTerm1.partialUpfront": "0.0844",
            "yrTerm3.partialUpfront": "0.0682"
          },
          "ondemand": "0.13"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.048",
            "yrTerm3.allUpfront": "0.0261",
            "yrTerm1.allUpfront": "0.0403",
            "yrTerm1.partialUpfront": "0.0411",
            "yrTerm3.partialUpfront": "0.0278"
          },
          "ondemand": "0.067"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.149",
            "yrTerm3.allUpfront": "0.0943",
            "yrTerm1.allUpfront": "0.1243",
            "yrTerm1.partialUpfront": "0.1268",
            "yrTerm3.partialUpfront": "0.1003"
          },
          "ondemand": "0.182"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.253",
            "yrTerm3.allUpfront": "0.1774",
            "yrTerm1.allUpfront": "0.2112",
            "yrTerm1.partialUpfront": "0.2151",
            "yrTerm3.partialUpfront": "0.1887"
          },
          "ondemand": "0.353"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.114",
            "yrTerm3.allUpfront": "0.0759",
            "yrTerm1.allUpfront": "0.0953",
            "yrTerm1.partialUpfront": "0.0971",
            "yrTerm3.partialUpfront": "0.0808"
          },
          "ondemand": "0.146"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.064",
            "yrTerm3.allUpfront": "0.0369",
            "yrTerm1.allUpfront": "0.0538",
            "yrTerm1.partialUpfront": "0.0545",
            "yrTerm3.partialUpfront": "0.0393"
          },
          "ondemand": "0.096"
        }
      },
      "sa-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.189",
            "yrTerm3.allUpfront": "0.1056",
            "yrTerm1.allUpfront": "0.1586",
            "yrTerm1.partialUpfront": "0.1618",
            "yrTerm3.partialUpfront": "0.1124"
          },
          "ondemand": "0.235"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.293",
            "yrTerm3.allUpfront": "0.1877",
            "yrTerm1.allUpfront": "0.2451",
            "yrTerm1.partialUpfront": "0.2501",
            "yrTerm3.partialUpfront": "0.1997"
          },
          "ondemand": "0.419"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.149",
            "yrTerm3.allUpfront": "0.0825",
            "yrTerm1.allUpfront": "0.141",
            "yrTerm1.partialUpfront": "0.1438",
            "yrTerm3.partialUpfront": "0.0877"
          },
          "ondemand": "0.158"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.06",
            "yrTerm3.allUpfront": "0.036",
            "yrTerm1.allUpfront": "0.0509",
            "yrTerm1.partialUpfront": "0.0521",
            "yrTerm3.partialUpfront": "0.0383"
          },
          "ondemand": "0.095"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.148",
            "yrTerm3.allUpfront": "0.0909",
            "yrTerm1.allUpfront": "0.1241",
            "yrTerm1.partialUpfront": "0.1266",
            "yrTerm3.partialUpfront": "0.0967"
          },
          "ondemand": "0.187"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.243",
            "yrTerm3.allUpfront": "0.1673",
            "yrTerm1.allUpfront": "0.2038",
            "yrTerm1.partialUpfront": "0.2079",
            "yrTerm3.partialUpfront": "0.178"
          },
          "ondemand": "0.356"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.116",
            "yrTerm3.allUpfront": "0.0725",
            "yrTerm1.allUpfront": "0.0973",
            "yrTerm1.partialUpfront": "0.0992",
            "yrTerm3.partialUpfront": "0.0772"
          },
          "ondemand": "0.161"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.068",
            "yrTerm3.allUpfront": "0.0367",
            "yrTerm1.allUpfront": "0.0578",
            "yrTerm1.partialUpfront": "0.059",
            "yrTerm3.partialUpfront": "0.0391"
          },
          "ondemand": "0.098"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.145",
            "yrTerm3.allUpfront": "0.0909",
            "yrTerm1.allUpfront": "0.1212",
            "yrTerm1.partialUpfront": "0.1246",
            "yrTerm3.partialUpfront": "0.0967"
          },
          "ondemand": "0.182"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.24",
            "yrTerm3.allUpfront": "0.1673",
            "yrTerm1.allUpfront": "0.2012",
            "yrTerm1.partialUpfront": "0.2048",
            "yrTerm3.partialUpfront": "0.178"
          },
          "ondemand": "0.351"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.113",
            "yrTerm3.allUpfront": "0.0725",
            "yrTerm1.allUpfront": "0.0941",
            "yrTerm1.partialUpfront": "0.0962",
            "yrTerm3.partialUpfront": "0.0772"
          },
          "ondemand": "0.156"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.065",
            "yrTerm3.allUpfront": "0.0367",
            "yrTerm1.allUpfront": "0.055",
            "yrTerm1.partialUpfront": "0.0556",
            "yrTerm3.partialUpfront": "0.0391"
          },
          "ondemand": "0.093"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.127",
            "yrTerm3.allUpfront": "0.0785",
            "yrTerm1.allUpfront": "0.1059",
            "yrTerm1.partialUpfront": "0.1082",
            "yrTerm3.partialUpfront": "0.0836"
          },
          "ondemand": "0.184"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.221",
            "yrTerm3.allUpfront": "0.155",
            "yrTerm1.allUpfront": "0.184",
            "yrTerm1.partialUpfront": "0.1883",
            "yrTerm3.partialUpfront": "0.1649"
          },
          "ondemand": "0.353"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.1",
            "yrTerm3.allUpfront": "0.0641",
            "yrTerm1.allUpfront": "0.084",
            "yrTerm1.partialUpfront": "0.0844",
            "yrTerm3.partialUpfront": "0.0682"
          },
          "ondemand": "0.13"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.048",
            "yrTerm3.allUpfront": "0.0261",
            "yrTerm1.allUpfront": "0.0403",
            "yrTerm1.partialUpfront": "0.0411",
            "yrTerm3.partialUpfront": "0.0278"
          },
          "ondemand": "0.067"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.179",
            "yrTerm3.allUpfront": "0.1095",
            "yrTerm1.allUpfront": "0.1395",
            "yrTerm1.partialUpfront": "0.1533",
            "yrTerm3.partialUpfront": "0.1203"
          },
          "ondemand": "0.222"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.11",
            "yrTerm3.allUpfront": "0.0683",
            "yrTerm1.allUpfront": "0.0863",
            "yrTerm1.partialUpfront": "0.0944",
            "yrTerm3.partialUpfront": "0.0765"
          },
          "ondemand": "0.147"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.061",
            "yrTerm3.allUpfront": "0.0322",
            "yrTerm1.allUpfront": "0.05",
            "yrTerm1.partialUpfront": "0.0527",
            "yrTerm3.partialUpfront": "0.0347"
          },
          "ondemand": "0.084"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.148",
            "yrTerm3.allUpfront": "0.0909",
            "yrTerm1.allUpfront": "0.1241",
            "yrTerm1.partialUpfront": "0.1266",
            "yrTerm3.partialUpfront": "0.0967"
          },
          "ondemand": "0.201"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.243",
            "yrTerm3.allUpfront": "0.1673",
            "yrTerm1.allUpfront": "0.2038",
            "yrTerm1.partialUpfront": "0.2079",
            "yrTerm3.partialUpfront": "0.178"
          },
          "ondemand": "0.374"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.117",
            "yrTerm3.allUpfront": "0.0735",
            "yrTerm1.allUpfront": "0.0983",
            "yrTerm1.partialUpfront": "0.1002",
            "yrTerm3.partialUpfront": "0.0782"
          },
          "ondemand": "0.14"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.058",
            "yrTerm3.allUpfront": "0.0336",
            "yrTerm1.allUpfront": "0.0532",
            "yrTerm1.partialUpfront": "0.0543",
            "yrTerm3.partialUpfront": "0.0358"
          },
          "ondemand": "0.077"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.143",
            "yrTerm3.allUpfront": "0.0884",
            "yrTerm1.allUpfront": "0.1192",
            "yrTerm1.partialUpfront": "0.1217",
            "yrTerm3.partialUpfront": "0.094"
          },
          "ondemand": "0.243"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.116",
            "yrTerm3.allUpfront": "0.0743",
            "yrTerm1.allUpfront": "0.0972",
            "yrTerm1.partialUpfront": "0.0989",
            "yrTerm3.partialUpfront": "0.079"
          },
          "ondemand": "0.142"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.06",
            "yrTerm3.allUpfront": "0.0363",
            "yrTerm1.allUpfront": "0.0536",
            "yrTerm1.partialUpfront": "0.0551",
            "yrTerm3.partialUpfront": "0.0387"
          },
          "ondemand": "0.079"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.137",
            "yrTerm3.allUpfront": "0.0859",
            "yrTerm1.allUpfront": "0.1144",
            "yrTerm1.partialUpfront": "0.1177",
            "yrTerm3.partialUpfront": "0.0914"
          },
          "ondemand": "0.183"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.22",
            "yrTerm3.allUpfront": "0.155",
            "yrTerm1.allUpfront": "0.1835",
            "yrTerm1.partialUpfront": "0.1872",
            "yrTerm3.partialUpfront": "0.1649"
          },
          "ondemand": "0.352"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.114",
            "yrTerm3.allUpfront": "0.0735",
            "yrTerm1.allUpfront": "0.0955",
            "yrTerm1.partialUpfront": "0.0973",
            "yrTerm3.partialUpfront": "0.0782"
          },
          "ondemand": "0.129"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.055",
            "yrTerm3.allUpfront": "0.0336",
            "yrTerm1.allUpfront": "0.0496",
            "yrTerm1.partialUpfront": "0.0511",
            "yrTerm3.partialUpfront": "0.0358"
          },
          "ondemand": "0.073"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 6,
      "max_enis": 2
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "HVM",
      "PV"
    ],
    "ebs_optimized": false,
    "storage": {
      "ssd": true,
      "devices": 1,
      "size": 4.0
    },
    "instance_type": "m3.medium",
    "ECU": 3.0,
    "memory": 3.75,
    "ebs_max_bandwidth": 0
  },
  {
    "family": "General purpose",
    "enhanced_networking": false,
    "vCPU": 2,
    "generation": "current",
    "ebs_iops": 0,
    "network_performance": "Moderate",
    "ebs_throughput": 0,
    "pretty_name": "M3 General Purpose Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.239",
            "yrTerm3.allUpfront": "0.1502",
            "yrTerm1.allUpfront": "0.1997",
            "yrTerm1.partialUpfront": "0.205",
            "yrTerm3.partialUpfront": "0.1598"
          },
          "ondemand": "0.367"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.442",
            "yrTerm3.allUpfront": "0.31",
            "yrTerm1.allUpfront": "0.3698",
            "yrTerm1.partialUpfront": "0.3781",
            "yrTerm3.partialUpfront": "0.3298"
          },
          "ondemand": "0.704"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.199",
            "yrTerm3.allUpfront": "0.1272",
            "yrTerm1.allUpfront": "0.1667",
            "yrTerm1.partialUpfront": "0.1702",
            "yrTerm3.partialUpfront": "0.1353"
          },
          "ondemand": "0.259"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.095",
            "yrTerm3.allUpfront": "0.0522",
            "yrTerm1.allUpfront": "0.0814",
            "yrTerm1.partialUpfront": "0.0831",
            "yrTerm3.partialUpfront": "0.0556"
          },
          "ondemand": "0.133"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.283",
            "yrTerm3.allUpfront": "0.18",
            "yrTerm1.allUpfront": "0.2374",
            "yrTerm1.partialUpfront": "0.2426",
            "yrTerm3.partialUpfront": "0.1915"
          },
          "ondemand": "0.364"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.504",
            "yrTerm3.allUpfront": "0.3557",
            "yrTerm1.allUpfront": "0.4223",
            "yrTerm1.partialUpfront": "0.4315",
            "yrTerm3.partialUpfront": "0.3784"
          },
          "ondemand": "0.706"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.228",
            "yrTerm3.allUpfront": "0.1498",
            "yrTerm1.allUpfront": "0.1914",
            "yrTerm1.partialUpfront": "0.1949",
            "yrTerm3.partialUpfront": "0.1593"
          },
          "ondemand": "0.292"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.127",
            "yrTerm3.allUpfront": "0.0738",
            "yrTerm1.allUpfront": "0.1085",
            "yrTerm1.partialUpfront": "0.1109",
            "yrTerm3.partialUpfront": "0.0785"
          },
          "ondemand": "0.193"
        }
      },
      "sa-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.358",
            "yrTerm3.allUpfront": "0.2012",
            "yrTerm1.allUpfront": "0.2998",
            "yrTerm1.partialUpfront": "0.3059",
            "yrTerm3.partialUpfront": "0.214"
          },
          "ondemand": "0.469"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.585",
            "yrTerm3.allUpfront": "0.3745",
            "yrTerm1.allUpfront": "0.4902",
            "yrTerm1.partialUpfront": "0.5002",
            "yrTerm3.partialUpfront": "0.3984"
          },
          "ondemand": "0.837"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.296",
            "yrTerm3.allUpfront": "0.1675",
            "yrTerm1.allUpfront": "0.2799",
            "yrTerm1.partialUpfront": "0.2857",
            "yrTerm3.partialUpfront": "0.1782"
          },
          "ondemand": "0.316"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.12",
            "yrTerm3.allUpfront": "0.0729",
            "yrTerm1.allUpfront": "0.1011",
            "yrTerm1.partialUpfront": "0.1032",
            "yrTerm3.partialUpfront": "0.0776"
          },
          "ondemand": "0.19"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.281",
            "yrTerm3.allUpfront": "0.1734",
            "yrTerm1.allUpfront": "0.2353",
            "yrTerm1.partialUpfront": "0.2401",
            "yrTerm3.partialUpfront": "0.1845"
          },
          "ondemand": "0.374"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.485",
            "yrTerm3.allUpfront": "0.3346",
            "yrTerm1.allUpfront": "0.4065",
            "yrTerm1.partialUpfront": "0.4148",
            "yrTerm3.partialUpfront": "0.356"
          },
          "ondemand": "0.711"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.23",
            "yrTerm3.allUpfront": "0.1441",
            "yrTerm1.allUpfront": "0.1928",
            "yrTerm1.partialUpfront": "0.1967",
            "yrTerm3.partialUpfront": "0.1533"
          },
          "ondemand": "0.322"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.136",
            "yrTerm3.allUpfront": "0.0735",
            "yrTerm1.allUpfront": "0.1156",
            "yrTerm1.partialUpfront": "0.1181",
            "yrTerm3.partialUpfront": "0.0782"
          },
          "ondemand": "0.196"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.274",
            "yrTerm3.allUpfront": "0.1734",
            "yrTerm1.allUpfront": "0.2293",
            "yrTerm1.partialUpfront": "0.2338",
            "yrTerm3.partialUpfront": "0.1845"
          },
          "ondemand": "0.364"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.478",
            "yrTerm3.allUpfront": "0.3346",
            "yrTerm1.allUpfront": "0.4013",
            "yrTerm1.partialUpfront": "0.4095",
            "yrTerm3.partialUpfront": "0.356"
          },
          "ondemand": "0.701"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.223",
            "yrTerm3.allUpfront": "0.1441",
            "yrTerm1.allUpfront": "0.1873",
            "yrTerm1.partialUpfront": "0.1908",
            "yrTerm3.partialUpfront": "0.1533"
          },
          "ondemand": "0.312"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.129",
            "yrTerm3.allUpfront": "0.0735",
            "yrTerm1.allUpfront": "0.1099",
            "yrTerm1.partialUpfront": "0.1122",
            "yrTerm3.partialUpfront": "0.0782"
          },
          "ondemand": "0.186"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.239",
            "yrTerm3.allUpfront": "0.1502",
            "yrTerm1.allUpfront": "0.1997",
            "yrTerm1.partialUpfront": "0.205",
            "yrTerm3.partialUpfront": "0.1598"
          },
          "ondemand": "0.367"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.442",
            "yrTerm3.allUpfront": "0.31",
            "yrTerm1.allUpfront": "0.3698",
            "yrTerm1.partialUpfront": "0.3781",
            "yrTerm3.partialUpfront": "0.3298"
          },
          "ondemand": "0.704"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.199",
            "yrTerm3.allUpfront": "0.1272",
            "yrTerm1.allUpfront": "0.1667",
            "yrTerm1.partialUpfront": "0.1702",
            "yrTerm3.partialUpfront": "0.1353"
          },
          "ondemand": "0.259"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.095",
            "yrTerm3.allUpfront": "0.0522",
            "yrTerm1.allUpfront": "0.0814",
            "yrTerm1.partialUpfront": "0.0831",
            "yrTerm3.partialUpfront": "0.0556"
          },
          "ondemand": "0.133"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.358",
            "yrTerm3.allUpfront": "0.2189",
            "yrTerm1.allUpfront": "0.2782",
            "yrTerm1.partialUpfront": "0.3056",
            "yrTerm3.partialUpfront": "0.2405"
          },
          "ondemand": "0.444"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.753",
            "yrTerm3.allUpfront": "0.4992",
            "yrTerm1.allUpfront": "0.578",
            "yrTerm1.partialUpfront": "0.6438",
            "yrTerm3.partialUpfront": "0.5561"
          },
          "ondemand": "0.798"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.221",
            "yrTerm3.allUpfront": "0.1357",
            "yrTerm1.allUpfront": "0.1725",
            "yrTerm1.partialUpfront": "0.1887",
            "yrTerm3.partialUpfront": "0.152"
          },
          "ondemand": "0.294"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.122",
            "yrTerm3.allUpfront": "0.0643",
            "yrTerm1.allUpfront": "0.1",
            "yrTerm1.partialUpfront": "0.1054",
            "yrTerm3.partialUpfront": "0.0694"
          },
          "ondemand": "0.168"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.281",
            "yrTerm3.allUpfront": "0.1734",
            "yrTerm1.allUpfront": "0.2353",
            "yrTerm1.partialUpfront": "0.2401",
            "yrTerm3.partialUpfront": "0.1845"
          },
          "ondemand": "0.401"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.485",
            "yrTerm3.allUpfront": "0.3346",
            "yrTerm1.allUpfront": "0.4065",
            "yrTerm1.partialUpfront": "0.4148",
            "yrTerm3.partialUpfront": "0.356"
          },
          "ondemand": "0.747"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.235",
            "yrTerm3.allUpfront": "0.146",
            "yrTerm1.allUpfront": "0.1967",
            "yrTerm1.partialUpfront": "0.2007",
            "yrTerm3.partialUpfront": "0.1553"
          },
          "ondemand": "0.28"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.116",
            "yrTerm3.allUpfront": "0.0673",
            "yrTerm1.allUpfront": "0.1054",
            "yrTerm1.partialUpfront": "0.1076",
            "yrTerm3.partialUpfront": "0.0716"
          },
          "ondemand": "0.154"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.271",
            "yrTerm3.allUpfront": "0.1705",
            "yrTerm1.allUpfront": "0.2264",
            "yrTerm1.partialUpfront": "0.2307",
            "yrTerm3.partialUpfront": "0.1813"
          },
          "ondemand": "0.466"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.858",
            "yrTerm3.allUpfront": "0.6124",
            "yrTerm1.allUpfront": "0.7174",
            "yrTerm1.partialUpfront": "0.7317",
            "yrTerm3.partialUpfront": "0.6515"
          },
          "ondemand": "1.25"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.231",
            "yrTerm3.allUpfront": "0.1477",
            "yrTerm1.allUpfront": "0.1933",
            "yrTerm1.partialUpfront": "0.1977",
            "yrTerm3.partialUpfront": "0.1571"
          },
          "ondemand": "0.284"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.12",
            "yrTerm3.allUpfront": "0.0726",
            "yrTerm1.allUpfront": "0.108",
            "yrTerm1.partialUpfront": "0.1102",
            "yrTerm3.partialUpfront": "0.0773"
          },
          "ondemand": "0.158"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.259",
            "yrTerm3.allUpfront": "0.1637",
            "yrTerm1.allUpfront": "0.2168",
            "yrTerm1.partialUpfront": "0.2202",
            "yrTerm3.partialUpfront": "0.1741"
          },
          "ondemand": "0.366"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.441",
            "yrTerm3.allUpfront": "0.31",
            "yrTerm1.allUpfront": "0.3688",
            "yrTerm1.partialUpfront": "0.3766",
            "yrTerm3.partialUpfront": "0.3298"
          },
          "ondemand": "0.703"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.229",
            "yrTerm3.allUpfront": "0.1469",
            "yrTerm1.allUpfront": "0.1918",
            "yrTerm1.partialUpfront": "0.1961",
            "yrTerm3.partialUpfront": "0.1563"
          },
          "ondemand": "0.258"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.11",
            "yrTerm3.allUpfront": "0.0673",
            "yrTerm1.allUpfront": "0.1002",
            "yrTerm1.partialUpfront": "0.1021",
            "yrTerm3.partialUpfront": "0.0716"
          },
          "ondemand": "0.146"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 10,
      "max_enis": 3
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "HVM",
      "PV"
    ],
    "ebs_optimized": false,
    "storage": {
      "ssd": true,
      "devices": 1,
      "size": 32.0
    },
    "instance_type": "m3.large",
    "ECU": 6.5,
    "memory": 7.5,
    "ebs_max_bandwidth": 0
  },
  {
    "family": "General purpose",
    "enhanced_networking": false,
    "vCPU": 4,
    "generation": "current",
    "ebs_iops": 62.5,
    "network_performance": "High",
    "ebs_throughput": 4000.0,
    "pretty_name": "M3 General Purpose Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.479",
            "yrTerm3.allUpfront": "0.2995",
            "yrTerm1.allUpfront": "0.4014",
            "yrTerm1.partialUpfront": "0.4096",
            "yrTerm3.partialUpfront": "0.3186"
          },
          "ondemand": "0.734"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.882",
            "yrTerm3.allUpfront": "0.619",
            "yrTerm1.allUpfront": "0.7385",
            "yrTerm1.partialUpfront": "0.7536",
            "yrTerm3.partialUpfront": "0.6585"
          },
          "ondemand": "1.266"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.397",
            "yrTerm3.allUpfront": "0.2553",
            "yrTerm1.allUpfront": "0.3325",
            "yrTerm1.partialUpfront": "0.3396",
            "yrTerm3.partialUpfront": "0.2715"
          },
          "ondemand": "0.518"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.19",
            "yrTerm3.allUpfront": "0.1045",
            "yrTerm1.allUpfront": "0.1631",
            "yrTerm1.partialUpfront": "0.1662",
            "yrTerm3.partialUpfront": "0.1112"
          },
          "ondemand": "0.266"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.564",
            "yrTerm3.allUpfront": "0.3592",
            "yrTerm1.allUpfront": "0.4725",
            "yrTerm1.partialUpfront": "0.4816",
            "yrTerm3.partialUpfront": "0.3821"
          },
          "ondemand": "0.728"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.009",
            "yrTerm3.allUpfront": "0.7104",
            "yrTerm1.allUpfront": "0.8446",
            "yrTerm1.partialUpfront": "0.8621",
            "yrTerm3.partialUpfront": "0.7558"
          },
          "ondemand": "1.333"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.457",
            "yrTerm3.allUpfront": "0.2984",
            "yrTerm1.allUpfront": "0.3816",
            "yrTerm1.partialUpfront": "0.3895",
            "yrTerm3.partialUpfront": "0.3174"
          },
          "ondemand": "0.583"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.255",
            "yrTerm3.allUpfront": "0.1467",
            "yrTerm1.allUpfront": "0.218",
            "yrTerm1.partialUpfront": "0.2225",
            "yrTerm3.partialUpfront": "0.1561"
          },
          "ondemand": "0.385"
        }
      },
      "sa-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.715",
            "yrTerm3.allUpfront": "0.4014",
            "yrTerm1.allUpfront": "0.5993",
            "yrTerm1.partialUpfront": "0.6115",
            "yrTerm3.partialUpfront": "0.4271"
          },
          "ondemand": "0.939"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.17",
            "yrTerm3.allUpfront": "0.749",
            "yrTerm1.allUpfront": "0.9804",
            "yrTerm1.partialUpfront": "1.0003",
            "yrTerm3.partialUpfront": "0.7968"
          },
          "ondemand": "1.507"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.593",
            "yrTerm3.allUpfront": "0.3334",
            "yrTerm1.allUpfront": "0.561",
            "yrTerm1.partialUpfront": "0.5724",
            "yrTerm3.partialUpfront": "0.3547"
          },
          "ondemand": "0.633"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.24",
            "yrTerm3.allUpfront": "0.1448",
            "yrTerm1.allUpfront": "0.2031",
            "yrTerm1.partialUpfront": "0.2073",
            "yrTerm3.partialUpfront": "0.1541"
          },
          "ondemand": "0.381"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.563",
            "yrTerm3.allUpfront": "0.3459",
            "yrTerm1.allUpfront": "0.4715",
            "yrTerm1.partialUpfront": "0.4811",
            "yrTerm3.partialUpfront": "0.368"
          },
          "ondemand": "0.748"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.971",
            "yrTerm3.allUpfront": "0.6683",
            "yrTerm1.allUpfront": "0.8129",
            "yrTerm1.partialUpfront": "0.8295",
            "yrTerm3.partialUpfront": "0.7109"
          },
          "ondemand": "1.28"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.46",
            "yrTerm3.allUpfront": "0.2881",
            "yrTerm1.allUpfront": "0.3852",
            "yrTerm1.partialUpfront": "0.393",
            "yrTerm3.partialUpfront": "0.3065"
          },
          "ondemand": "0.644"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.272",
            "yrTerm3.allUpfront": "0.147",
            "yrTerm1.allUpfront": "0.2312",
            "yrTerm1.partialUpfront": "0.236",
            "yrTerm3.partialUpfront": "0.1564"
          },
          "ondemand": "0.392"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.549",
            "yrTerm3.allUpfront": "0.3459",
            "yrTerm1.allUpfront": "0.4595",
            "yrTerm1.partialUpfront": "0.4694",
            "yrTerm3.partialUpfront": "0.368"
          },
          "ondemand": "0.728"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.957",
            "yrTerm3.allUpfront": "0.6683",
            "yrTerm1.allUpfront": "0.8015",
            "yrTerm1.partialUpfront": "0.818",
            "yrTerm3.partialUpfront": "0.7109"
          },
          "ondemand": "1.26"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.446",
            "yrTerm3.allUpfront": "0.2881",
            "yrTerm1.allUpfront": "0.3735",
            "yrTerm1.partialUpfront": "0.3809",
            "yrTerm3.partialUpfront": "0.3065"
          },
          "ondemand": "0.624"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.258",
            "yrTerm3.allUpfront": "0.147",
            "yrTerm1.allUpfront": "0.2197",
            "yrTerm1.partialUpfront": "0.2244",
            "yrTerm3.partialUpfront": "0.1564"
          },
          "ondemand": "0.372"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.479",
            "yrTerm3.allUpfront": "0.2995",
            "yrTerm1.allUpfront": "0.4014",
            "yrTerm1.partialUpfront": "0.4096",
            "yrTerm3.partialUpfront": "0.3186"
          },
          "ondemand": "0.734"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.882",
            "yrTerm3.allUpfront": "0.619",
            "yrTerm1.allUpfront": "0.7385",
            "yrTerm1.partialUpfront": "0.7536",
            "yrTerm3.partialUpfront": "0.6585"
          },
          "ondemand": "1.266"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.397",
            "yrTerm3.allUpfront": "0.2553",
            "yrTerm1.allUpfront": "0.3325",
            "yrTerm1.partialUpfront": "0.3396",
            "yrTerm3.partialUpfront": "0.2715"
          },
          "ondemand": "0.518"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.19",
            "yrTerm3.allUpfront": "0.1045",
            "yrTerm1.allUpfront": "0.1631",
            "yrTerm1.partialUpfront": "0.1662",
            "yrTerm3.partialUpfront": "0.1112"
          },
          "ondemand": "0.266"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.511",
            "yrTerm3.allUpfront": "0.313",
            "yrTerm1.allUpfront": "0.3978",
            "yrTerm1.partialUpfront": "0.4371",
            "yrTerm3.partialUpfront": "0.3439"
          },
          "ondemand": "0.887"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.942",
            "yrTerm3.allUpfront": "0.6242",
            "yrTerm1.allUpfront": "0.7227",
            "yrTerm1.partialUpfront": "0.8051",
            "yrTerm3.partialUpfront": "0.6953"
          },
          "ondemand": "1.436"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.441",
            "yrTerm3.allUpfront": "0.2704",
            "yrTerm1.allUpfront": "0.3449",
            "yrTerm1.partialUpfront": "0.3773",
            "yrTerm3.partialUpfront": "0.303"
          },
          "ondemand": "0.588"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.244",
            "yrTerm3.allUpfront": "0.1277",
            "yrTerm1.allUpfront": "0.1991",
            "yrTerm1.partialUpfront": "0.2097",
            "yrTerm3.partialUpfront": "0.1378"
          },
          "ondemand": "0.336"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.563",
            "yrTerm3.allUpfront": "0.3459",
            "yrTerm1.allUpfront": "0.4715",
            "yrTerm1.partialUpfront": "0.4811",
            "yrTerm3.partialUpfront": "0.368"
          },
          "ondemand": "0.803"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.971",
            "yrTerm3.allUpfront": "0.6683",
            "yrTerm1.allUpfront": "0.8129",
            "yrTerm1.partialUpfront": "0.8295",
            "yrTerm3.partialUpfront": "0.7109"
          },
          "ondemand": "1.345"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.467",
            "yrTerm3.allUpfront": "0.2928",
            "yrTerm1.allUpfront": "0.391",
            "yrTerm1.partialUpfront": "0.399",
            "yrTerm3.partialUpfront": "0.3115"
          },
          "ondemand": "0.56"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.232",
            "yrTerm3.allUpfront": "0.1346",
            "yrTerm1.allUpfront": "0.2098",
            "yrTerm1.partialUpfront": "0.2141",
            "yrTerm3.partialUpfront": "0.1432"
          },
          "ondemand": "0.308"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.542",
            "yrTerm3.allUpfront": "0.3393",
            "yrTerm1.allUpfront": "0.4537",
            "yrTerm1.partialUpfront": "0.4632",
            "yrTerm3.partialUpfront": "0.3609"
          },
          "ondemand": "0.882"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.259",
            "yrTerm3.allUpfront": "0.8998",
            "yrTerm1.allUpfront": "1.0547",
            "yrTerm1.partialUpfront": "1.0761",
            "yrTerm3.partialUpfront": "0.9572"
          },
          "ondemand": "1.778"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.46",
            "yrTerm3.allUpfront": "0.295",
            "yrTerm1.allUpfront": "0.3848",
            "yrTerm1.partialUpfront": "0.3926",
            "yrTerm3.partialUpfront": "0.3138"
          },
          "ondemand": "0.567"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.239",
            "yrTerm3.allUpfront": "0.1443",
            "yrTerm1.allUpfront": "0.2153",
            "yrTerm1.partialUpfront": "0.2194",
            "yrTerm3.partialUpfront": "0.1535"
          },
          "ondemand": "0.315"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.518",
            "yrTerm3.allUpfront": "0.3255",
            "yrTerm1.allUpfront": "0.4336",
            "yrTerm1.partialUpfront": "0.4428",
            "yrTerm3.partialUpfront": "0.3463"
          },
          "ondemand": "0.733"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.88",
            "yrTerm3.allUpfront": "0.619",
            "yrTerm1.allUpfront": "0.7366",
            "yrTerm1.partialUpfront": "0.7508",
            "yrTerm3.partialUpfront": "0.6585"
          },
          "ondemand": "1.265"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.455",
            "yrTerm3.allUpfront": "0.2928",
            "yrTerm1.allUpfront": "0.3805",
            "yrTerm1.partialUpfront": "0.3884",
            "yrTerm3.partialUpfront": "0.3115"
          },
          "ondemand": "0.517"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.22",
            "yrTerm3.allUpfront": "0.1336",
            "yrTerm1.allUpfront": "0.1994",
            "yrTerm1.partialUpfront": "0.2032",
            "yrTerm3.partialUpfront": "0.1422"
          },
          "ondemand": "0.293"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 15,
      "max_enis": 4
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "HVM",
      "PV"
    ],
    "ebs_optimized": true,
    "storage": {
      "ssd": true,
      "devices": 2,
      "size": 40.0
    },
    "instance_type": "m3.xlarge",
    "ECU": 13.0,
    "memory": 15.0,
    "ebs_max_bandwidth": 500.0
  },
  {
    "family": "General purpose",
    "enhanced_networking": false,
    "vCPU": 8,
    "generation": "current",
    "ebs_iops": 125.0,
    "network_performance": "High",
    "ebs_throughput": 8000.0,
    "pretty_name": "M3 General Purpose Double Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.957",
            "yrTerm3.allUpfront": "0.5998",
            "yrTerm1.allUpfront": "0.802",
            "yrTerm1.partialUpfront": "0.8173",
            "yrTerm3.partialUpfront": "0.6381"
          },
          "ondemand": "1.468"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.661",
            "yrTerm3.allUpfront": "1.1566",
            "yrTerm1.allUpfront": "1.391",
            "yrTerm1.partialUpfront": "1.42",
            "yrTerm3.partialUpfront": "1.2304"
          },
          "ondemand": "2.532"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.793",
            "yrTerm3.allUpfront": "0.5105",
            "yrTerm1.allUpfront": "0.664",
            "yrTerm1.partialUpfront": "0.6781",
            "yrTerm3.partialUpfront": "0.5431"
          },
          "ondemand": "1.036"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.38",
            "yrTerm3.allUpfront": "0.209",
            "yrTerm1.allUpfront": "0.3243",
            "yrTerm1.partialUpfront": "0.3312",
            "yrTerm3.partialUpfront": "0.2224"
          },
          "ondemand": "0.532"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.128",
            "yrTerm3.allUpfront": "0.7192",
            "yrTerm1.allUpfront": "0.9451",
            "yrTerm1.partialUpfront": "0.9639",
            "yrTerm3.partialUpfront": "0.7651"
          },
          "ondemand": "1.456"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.854",
            "yrTerm3.allUpfront": "1.3078",
            "yrTerm1.allUpfront": "1.5531",
            "yrTerm1.partialUpfront": "1.5845",
            "yrTerm3.partialUpfront": "1.3913"
          },
          "ondemand": "2.665"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.91",
            "yrTerm3.allUpfront": "0.5951",
            "yrTerm1.allUpfront": "0.7622",
            "yrTerm1.partialUpfront": "0.778",
            "yrTerm3.partialUpfront": "0.6331"
          },
          "ondemand": "1.166"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.509",
            "yrTerm3.allUpfront": "0.2943",
            "yrTerm1.allUpfront": "0.436",
            "yrTerm1.partialUpfront": "0.4454",
            "yrTerm3.partialUpfront": "0.3132"
          },
          "ondemand": "0.77"
        }
      },
      "sa-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.431",
            "yrTerm3.allUpfront": "0.8038",
            "yrTerm1.allUpfront": "1.1987",
            "yrTerm1.partialUpfront": "1.2232",
            "yrTerm3.partialUpfront": "0.8552"
          },
          "ondemand": "1.877"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.341",
            "yrTerm3.allUpfront": "1.497",
            "yrTerm1.allUpfront": "1.9606",
            "yrTerm1.partialUpfront": "2.0007",
            "yrTerm3.partialUpfront": "1.5925"
          },
          "ondemand": "3.013"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.184",
            "yrTerm3.allUpfront": "0.6682",
            "yrTerm1.allUpfront": "1.1205",
            "yrTerm1.partialUpfront": "1.1434",
            "yrTerm3.partialUpfront": "0.7108"
          },
          "ondemand": "1.265"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.48",
            "yrTerm3.allUpfront": "0.2887",
            "yrTerm1.allUpfront": "0.4063",
            "yrTerm1.partialUpfront": "0.4146",
            "yrTerm3.partialUpfront": "0.3072"
          },
          "ondemand": "0.761"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.127",
            "yrTerm3.allUpfront": "0.6925",
            "yrTerm1.allUpfront": "0.9439",
            "yrTerm1.partialUpfront": "0.9632",
            "yrTerm3.partialUpfront": "0.7367"
          },
          "ondemand": "1.496"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.785",
            "yrTerm3.allUpfront": "1.2296",
            "yrTerm1.allUpfront": "1.4953",
            "yrTerm1.partialUpfront": "1.5258",
            "yrTerm3.partialUpfront": "1.3081"
          },
          "ondemand": "2.56"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.921",
            "yrTerm3.allUpfront": "0.5754",
            "yrTerm1.allUpfront": "0.7711",
            "yrTerm1.partialUpfront": "0.7869",
            "yrTerm3.partialUpfront": "0.6121"
          },
          "ondemand": "1.288"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.544",
            "yrTerm3.allUpfront": "0.2939",
            "yrTerm1.allUpfront": "0.4623",
            "yrTerm1.partialUpfront": "0.4719",
            "yrTerm3.partialUpfront": "0.3127"
          },
          "ondemand": "0.784"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.1",
            "yrTerm3.allUpfront": "0.6925",
            "yrTerm1.allUpfront": "0.9209",
            "yrTerm1.partialUpfront": "0.9397",
            "yrTerm3.partialUpfront": "0.7367"
          },
          "ondemand": "1.457"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.758",
            "yrTerm3.allUpfront": "1.2296",
            "yrTerm1.allUpfront": "1.472",
            "yrTerm1.partialUpfront": "1.5025",
            "yrTerm3.partialUpfront": "1.3081"
          },
          "ondemand": "2.521"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.894",
            "yrTerm3.allUpfront": "0.5754",
            "yrTerm1.allUpfront": "0.7479",
            "yrTerm1.partialUpfront": "0.7638",
            "yrTerm3.partialUpfront": "0.6121"
          },
          "ondemand": "1.249"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.517",
            "yrTerm3.allUpfront": "0.2939",
            "yrTerm1.allUpfront": "0.4393",
            "yrTerm1.partialUpfront": "0.4487",
            "yrTerm3.partialUpfront": "0.3127"
          },
          "ondemand": "0.745"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.957",
            "yrTerm3.allUpfront": "0.5998",
            "yrTerm1.allUpfront": "0.802",
            "yrTerm1.partialUpfront": "0.8173",
            "yrTerm3.partialUpfront": "0.6381"
          },
          "ondemand": "1.468"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.661",
            "yrTerm3.allUpfront": "1.1566",
            "yrTerm1.allUpfront": "1.391",
            "yrTerm1.partialUpfront": "1.42",
            "yrTerm3.partialUpfront": "1.2304"
          },
          "ondemand": "2.532"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.793",
            "yrTerm3.allUpfront": "0.5105",
            "yrTerm1.allUpfront": "0.664",
            "yrTerm1.partialUpfront": "0.6781",
            "yrTerm3.partialUpfront": "0.5431"
          },
          "ondemand": "1.036"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.38",
            "yrTerm3.allUpfront": "0.209",
            "yrTerm1.allUpfront": "0.3243",
            "yrTerm1.partialUpfront": "0.3312",
            "yrTerm3.partialUpfront": "0.2224"
          },
          "ondemand": "0.532"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.023",
            "yrTerm3.allUpfront": "0.6274",
            "yrTerm1.allUpfront": "0.7955",
            "yrTerm1.partialUpfront": "0.8741",
            "yrTerm3.partialUpfront": "0.6895"
          },
          "ondemand": "1.774"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.733",
            "yrTerm3.allUpfront": "1.1493",
            "yrTerm1.allUpfront": "1.3296",
            "yrTerm1.partialUpfront": "1.4809",
            "yrTerm3.partialUpfront": "1.2803"
          },
          "ondemand": "2.871"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.883",
            "yrTerm3.allUpfront": "0.5416",
            "yrTerm1.allUpfront": "0.6898",
            "yrTerm1.partialUpfront": "0.7546",
            "yrTerm3.partialUpfront": "0.6068"
          },
          "ondemand": "1.176"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.488",
            "yrTerm3.allUpfront": "0.2546",
            "yrTerm1.allUpfront": "0.3989",
            "yrTerm1.partialUpfront": "0.4203",
            "yrTerm3.partialUpfront": "0.2746"
          },
          "ondemand": "0.672"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.127",
            "yrTerm3.allUpfront": "0.6925",
            "yrTerm1.allUpfront": "0.9439",
            "yrTerm1.partialUpfront": "0.9632",
            "yrTerm3.partialUpfront": "0.7367"
          },
          "ondemand": "1.606"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.785",
            "yrTerm3.allUpfront": "1.2296",
            "yrTerm1.allUpfront": "1.4953",
            "yrTerm1.partialUpfront": "1.5258",
            "yrTerm3.partialUpfront": "1.3081"
          },
          "ondemand": "2.689"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.935",
            "yrTerm3.allUpfront": "0.5829",
            "yrTerm1.allUpfront": "0.7829",
            "yrTerm1.partialUpfront": "0.7989",
            "yrTerm3.partialUpfront": "0.6201"
          },
          "ondemand": "1.12"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.464",
            "yrTerm3.allUpfront": "0.2683",
            "yrTerm1.allUpfront": "0.4216",
            "yrTerm1.partialUpfront": "0.4303",
            "yrTerm3.partialUpfront": "0.2854"
          },
          "ondemand": "0.616"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.085",
            "yrTerm3.allUpfront": "0.6798",
            "yrTerm1.allUpfront": "0.9075",
            "yrTerm1.partialUpfront": "0.9256",
            "yrTerm3.partialUpfront": "0.7232"
          },
          "ondemand": "1.829"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.506",
            "yrTerm3.allUpfront": "1.7911",
            "yrTerm1.allUpfront": "2.0974",
            "yrTerm1.partialUpfront": "2.1401",
            "yrTerm3.partialUpfront": "1.9054"
          },
          "ondemand": "3.537"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.921",
            "yrTerm3.allUpfront": "0.5901",
            "yrTerm1.allUpfront": "0.7695",
            "yrTerm1.partialUpfront": "0.7853",
            "yrTerm3.partialUpfront": "0.6278"
          },
          "ondemand": "1.136"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.479",
            "yrTerm3.allUpfront": "0.2887",
            "yrTerm1.allUpfront": "0.4296",
            "yrTerm1.partialUpfront": "0.4386",
            "yrTerm3.partialUpfront": "0.3071"
          },
          "ondemand": "0.632"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.037",
            "yrTerm3.allUpfront": "0.6518",
            "yrTerm1.allUpfront": "0.8661",
            "yrTerm1.partialUpfront": "0.8843",
            "yrTerm3.partialUpfront": "0.6934"
          },
          "ondemand": "1.465"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.658",
            "yrTerm3.allUpfront": "1.1566",
            "yrTerm1.allUpfront": "1.3872",
            "yrTerm1.partialUpfront": "1.4155",
            "yrTerm3.partialUpfront": "1.2304"
          },
          "ondemand": "2.529"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.912",
            "yrTerm3.allUpfront": "0.5839",
            "yrTerm1.allUpfront": "0.7622",
            "yrTerm1.partialUpfront": "0.7776",
            "yrTerm3.partialUpfront": "0.6211"
          },
          "ondemand": "1.033"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.441",
            "yrTerm3.allUpfront": "0.2673",
            "yrTerm1.allUpfront": "0.3979",
            "yrTerm1.partialUpfront": "0.4062",
            "yrTerm3.partialUpfront": "0.2844"
          },
          "ondemand": "0.585"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 30,
      "max_enis": 4
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "HVM",
      "PV"
    ],
    "ebs_optimized": true,
    "storage": {
      "ssd": true,
      "devices": 2,
      "size": 80.0
    },
    "instance_type": "m3.2xlarge",
    "ECU": 26.0,
    "memory": 30.0,
    "ebs_max_bandwidth": 1000.0
  },
  {
    "family": "Compute optimized",
    "enhanced_networking": true,
    "vCPU": 2,
    "generation": "current",
    "ebs_iops": 62.5,
    "network_performance": "Moderate",
    "ebs_throughput": 4000.0,
    "pretty_name": "C4 High-CPU Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.313",
            "yrTerm3.allUpfront": "0.1857",
            "yrTerm1.allUpfront": "0.262",
            "yrTerm1.partialUpfront": "0.2664",
            "yrTerm3.partialUpfront": "0.1977"
          },
          "ondemand": "0.413"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.068",
            "yrTerm3.allUpfront": "0.652",
            "yrTerm1.allUpfront": "0.8939",
            "yrTerm1.partialUpfront": "0.9122",
            "yrTerm3.partialUpfront": "0.6937"
          },
          "ondemand": "1.405"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.17",
            "yrTerm3.allUpfront": "0.1126",
            "yrTerm1.allUpfront": "0.1453",
            "yrTerm1.partialUpfront": "0.1481",
            "yrTerm3.partialUpfront": "0.1198"
          },
          "ondemand": "0.193"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.078",
            "yrTerm3.allUpfront": "0.0408",
            "yrTerm1.allUpfront": "0.0658",
            "yrTerm1.partialUpfront": "0.0667",
            "yrTerm3.partialUpfront": "0.0436"
          },
          "ondemand": "0.105"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.34",
            "yrTerm3.allUpfront": "0.2089",
            "yrTerm1.allUpfront": "0.2844",
            "yrTerm1.partialUpfront": "0.2902",
            "yrTerm3.partialUpfront": "0.222"
          },
          "ondemand": "0.441"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.095",
            "yrTerm3.allUpfront": "0.6752",
            "yrTerm1.allUpfront": "0.9174",
            "yrTerm1.partialUpfront": "0.9369",
            "yrTerm3.partialUpfront": "0.718"
          },
          "ondemand": "1.433"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.197",
            "yrTerm3.allUpfront": "0.137",
            "yrTerm1.allUpfront": "0.1691",
            "yrTerm1.partialUpfront": "0.1734",
            "yrTerm3.partialUpfront": "0.1457"
          },
          "ondemand": "0.233"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.101",
            "yrTerm3.allUpfront": "0.0639",
            "yrTerm1.allUpfront": "0.0888",
            "yrTerm1.partialUpfront": "0.0905",
            "yrTerm3.partialUpfront": "0.0678"
          },
          "ondemand": "0.133"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.2638",
            "yrTerm3.allUpfront": "0.2265",
            "yrTerm1.allUpfront": "0.2484",
            "yrTerm1.partialUpfront": "0.25",
            "yrTerm3.partialUpfront": "0.2302"
          },
          "ondemand": "0.289"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.9238",
            "yrTerm3.allUpfront": "0.8865",
            "yrTerm1.allUpfront": "0.9084",
            "yrTerm1.partialUpfront": "0.91",
            "yrTerm3.partialUpfront": "0.8902"
          },
          "ondemand": "0.949"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.1868",
            "yrTerm3.allUpfront": "0.1495",
            "yrTerm1.allUpfront": "0.1714",
            "yrTerm1.partialUpfront": "0.173",
            "yrTerm3.partialUpfront": "0.1532"
          },
          "ondemand": "0.212"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.0948",
            "yrTerm3.allUpfront": "0.0575",
            "yrTerm1.allUpfront": "0.0794",
            "yrTerm1.partialUpfront": "0.081",
            "yrTerm3.partialUpfront": "0.0612"
          },
          "ondemand": "0.12"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.342",
            "yrTerm3.allUpfront": "0.2056",
            "yrTerm1.allUpfront": "0.2871",
            "yrTerm1.partialUpfront": "0.2935",
            "yrTerm3.partialUpfront": "0.2184"
          },
          "ondemand": "0.452"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.097",
            "yrTerm3.allUpfront": "0.6718",
            "yrTerm1.allUpfront": "0.9187",
            "yrTerm1.partialUpfront": "0.9379",
            "yrTerm3.partialUpfront": "0.7144"
          },
          "ondemand": "1.444"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.199",
            "yrTerm3.allUpfront": "0.1335",
            "yrTerm1.allUpfront": "0.172",
            "yrTerm1.partialUpfront": "0.1756",
            "yrTerm3.partialUpfront": "0.142"
          },
          "ondemand": "0.244"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.107",
            "yrTerm3.allUpfront": "0.0609",
            "yrTerm1.allUpfront": "0.0909",
            "yrTerm1.partialUpfront": "0.0932",
            "yrTerm3.partialUpfront": "0.0645"
          },
          "ondemand": "0.144"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.337",
            "yrTerm3.allUpfront": "0.2056",
            "yrTerm1.allUpfront": "0.2826",
            "yrTerm1.partialUpfront": "0.2892",
            "yrTerm3.partialUpfront": "0.2184"
          },
          "ondemand": "0.445"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.092",
            "yrTerm3.allUpfront": "0.6718",
            "yrTerm1.allUpfront": "0.9146",
            "yrTerm1.partialUpfront": "0.9324",
            "yrTerm3.partialUpfront": "0.7144"
          },
          "ondemand": "1.437"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.194",
            "yrTerm3.allUpfront": "0.1335",
            "yrTerm1.allUpfront": "0.1672",
            "yrTerm1.partialUpfront": "0.1714",
            "yrTerm3.partialUpfront": "0.142"
          },
          "ondemand": "0.237"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.102",
            "yrTerm3.allUpfront": "0.0609",
            "yrTerm1.allUpfront": "0.0863",
            "yrTerm1.partialUpfront": "0.0889",
            "yrTerm3.partialUpfront": "0.0645"
          },
          "ondemand": "0.137"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.313",
            "yrTerm3.allUpfront": "0.1857",
            "yrTerm1.allUpfront": "0.262",
            "yrTerm1.partialUpfront": "0.2664",
            "yrTerm3.partialUpfront": "0.1977"
          },
          "ondemand": "0.413"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.068",
            "yrTerm3.allUpfront": "0.652",
            "yrTerm1.allUpfront": "0.8939",
            "yrTerm1.partialUpfront": "0.9122",
            "yrTerm3.partialUpfront": "0.6937"
          },
          "ondemand": "1.405"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.17",
            "yrTerm3.allUpfront": "0.1126",
            "yrTerm1.allUpfront": "0.1453",
            "yrTerm1.partialUpfront": "0.1481",
            "yrTerm3.partialUpfront": "0.1198"
          },
          "ondemand": "0.193"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.078",
            "yrTerm3.allUpfront": "0.0408",
            "yrTerm1.allUpfront": "0.0658",
            "yrTerm1.partialUpfront": "0.0667",
            "yrTerm3.partialUpfront": "0.0436"
          },
          "ondemand": "0.105"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.334",
            "yrTerm3.allUpfront": "0.2018",
            "yrTerm1.allUpfront": "0.2798",
            "yrTerm1.partialUpfront": "0.2858",
            "yrTerm3.partialUpfront": "0.2143"
          },
          "ondemand": "0.439"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.089",
            "yrTerm3.allUpfront": "0.668",
            "yrTerm1.allUpfront": "0.9123",
            "yrTerm1.partialUpfront": "0.9312",
            "yrTerm3.partialUpfront": "0.7103"
          },
          "ondemand": "1.431"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.191",
            "yrTerm3.allUpfront": "0.1296",
            "yrTerm1.allUpfront": "0.1643",
            "yrTerm1.partialUpfront": "0.1677",
            "yrTerm3.partialUpfront": "0.1378"
          },
          "ondemand": "0.221"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.098",
            "yrTerm3.allUpfront": "0.057",
            "yrTerm1.allUpfront": "0.0837",
            "yrTerm1.partialUpfront": "0.0857",
            "yrTerm3.partialUpfront": "0.0602"
          },
          "ondemand": "0.131"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.337",
            "yrTerm3.allUpfront": "0.2071",
            "yrTerm1.allUpfront": "0.2826",
            "yrTerm1.partialUpfront": "0.2892",
            "yrTerm3.partialUpfront": "0.22"
          },
          "ondemand": "0.442"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.091",
            "yrTerm3.allUpfront": "0.6733",
            "yrTerm1.allUpfront": "0.9137",
            "yrTerm1.partialUpfront": "0.9324",
            "yrTerm3.partialUpfront": "0.7161"
          },
          "ondemand": "1.434"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.194",
            "yrTerm3.allUpfront": "0.1351",
            "yrTerm1.allUpfront": "0.1672",
            "yrTerm1.partialUpfront": "0.1698",
            "yrTerm3.partialUpfront": "0.1437"
          },
          "ondemand": "0.224"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.1",
            "yrTerm3.allUpfront": "0.0621",
            "yrTerm1.allUpfront": "0.0862",
            "yrTerm1.partialUpfront": "0.0888",
            "yrTerm3.partialUpfront": "0.0658"
          },
          "ondemand": "0.134"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.325",
            "yrTerm3.allUpfront": "0.1985",
            "yrTerm1.allUpfront": "0.2718",
            "yrTerm1.partialUpfront": "0.2778",
            "yrTerm3.partialUpfront": "0.2117"
          },
          "ondemand": "0.427"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.08",
            "yrTerm3.allUpfront": "0.6648",
            "yrTerm1.allUpfront": "0.9048",
            "yrTerm1.partialUpfront": "0.9235",
            "yrTerm3.partialUpfront": "0.7077"
          },
          "ondemand": "1.419"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.182",
            "yrTerm3.allUpfront": "0.1261",
            "yrTerm1.allUpfront": "0.1567",
            "yrTerm1.partialUpfront": "0.1597",
            "yrTerm3.partialUpfront": "0.1342"
          },
          "ondemand": "0.208"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.089",
            "yrTerm3.allUpfront": "0.054",
            "yrTerm1.allUpfront": "0.0761",
            "yrTerm1.partialUpfront": "0.0778",
            "yrTerm3.partialUpfront": "0.0579"
          },
          "ondemand": "0.119"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 10,
      "max_enis": 3
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": true,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": true,
    "storage": null,
    "instance_type": "c4.large",
    "ECU": 8.0,
    "memory": 3.75,
    "ebs_max_bandwidth": 500.0
  },
  {
    "family": "Compute optimized",
    "enhanced_networking": true,
    "vCPU": 4,
    "generation": "current",
    "ebs_iops": 93.75,
    "network_performance": "High",
    "ebs_throughput": 6000.0,
    "pretty_name": "C4 High-CPU Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.589",
            "yrTerm3.allUpfront": "0.3489",
            "yrTerm1.allUpfront": "0.4937",
            "yrTerm1.partialUpfront": "0.5037",
            "yrTerm3.partialUpfront": "0.3714"
          },
          "ondemand": "0.776"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.27",
            "yrTerm3.allUpfront": "0.77",
            "yrTerm1.allUpfront": "1.0637",
            "yrTerm1.partialUpfront": "1.0858",
            "yrTerm3.partialUpfront": "0.8194"
          },
          "ondemand": "1.672"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.339",
            "yrTerm3.allUpfront": "0.226",
            "yrTerm1.allUpfront": "0.2887",
            "yrTerm1.partialUpfront": "0.2944",
            "yrTerm3.partialUpfront": "0.2404"
          },
          "ondemand": "0.386"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.155",
            "yrTerm3.allUpfront": "0.0816",
            "yrTerm1.allUpfront": "0.1314",
            "yrTerm1.partialUpfront": "0.1344",
            "yrTerm3.partialUpfront": "0.0872"
          },
          "ondemand": "0.209"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.645",
            "yrTerm3.allUpfront": "0.3952",
            "yrTerm1.allUpfront": "0.5398",
            "yrTerm1.partialUpfront": "0.5502",
            "yrTerm3.partialUpfront": "0.42"
          },
          "ondemand": "0.832"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.326",
            "yrTerm3.allUpfront": "0.8163",
            "yrTerm1.allUpfront": "1.1098",
            "yrTerm1.partialUpfront": "1.1323",
            "yrTerm3.partialUpfront": "0.8681"
          },
          "ondemand": "1.728"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.393",
            "yrTerm3.allUpfront": "0.2747",
            "yrTerm1.allUpfront": "0.3372",
            "yrTerm1.partialUpfront": "0.3441",
            "yrTerm3.partialUpfront": "0.2923"
          },
          "ondemand": "0.464"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.202",
            "yrTerm3.allUpfront": "0.1278",
            "yrTerm1.allUpfront": "0.1776",
            "yrTerm1.partialUpfront": "0.1811",
            "yrTerm3.partialUpfront": "0.1356"
          },
          "ondemand": "0.265"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.4476",
            "yrTerm3.allUpfront": "0.3721",
            "yrTerm1.allUpfront": "0.4166",
            "yrTerm1.partialUpfront": "0.4199",
            "yrTerm3.partialUpfront": "0.3794"
          },
          "ondemand": "0.496"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.0826",
            "yrTerm3.allUpfront": "1.0071",
            "yrTerm1.allUpfront": "1.0516",
            "yrTerm1.partialUpfront": "1.0549",
            "yrTerm3.partialUpfront": "1.0144"
          },
          "ondemand": "1.131"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.3746",
            "yrTerm3.allUpfront": "0.2991",
            "yrTerm1.allUpfront": "0.3436",
            "yrTerm1.partialUpfront": "0.3469",
            "yrTerm3.partialUpfront": "0.3064"
          },
          "ondemand": "0.423"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.1906",
            "yrTerm3.allUpfront": "0.1151",
            "yrTerm1.allUpfront": "0.1596",
            "yrTerm1.partialUpfront": "0.1629",
            "yrTerm3.partialUpfront": "0.1224"
          },
          "ondemand": "0.239"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.649",
            "yrTerm3.allUpfront": "0.3886",
            "yrTerm1.allUpfront": "0.5433",
            "yrTerm1.partialUpfront": "0.5543",
            "yrTerm3.partialUpfront": "0.4138"
          },
          "ondemand": "0.856"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.33",
            "yrTerm3.allUpfront": "0.8097",
            "yrTerm1.allUpfront": "1.1139",
            "yrTerm1.partialUpfront": "1.1365",
            "yrTerm3.partialUpfront": "0.8618"
          },
          "ondemand": "1.752"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.399",
            "yrTerm3.allUpfront": "0.2688",
            "yrTerm1.allUpfront": "0.3414",
            "yrTerm1.partialUpfront": "0.3484",
            "yrTerm3.partialUpfront": "0.2859"
          },
          "ondemand": "0.488"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.215",
            "yrTerm3.allUpfront": "0.1218",
            "yrTerm1.allUpfront": "0.1817",
            "yrTerm1.partialUpfront": "0.1854",
            "yrTerm3.partialUpfront": "0.13"
          },
          "ondemand": "0.289"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.638",
            "yrTerm3.allUpfront": "0.3886",
            "yrTerm1.allUpfront": "0.5342",
            "yrTerm1.partialUpfront": "0.5446",
            "yrTerm3.partialUpfront": "0.4138"
          },
          "ondemand": "0.842"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.319",
            "yrTerm3.allUpfront": "0.8097",
            "yrTerm1.allUpfront": "1.1053",
            "yrTerm1.partialUpfront": "1.1267",
            "yrTerm3.partialUpfront": "0.8618"
          },
          "ondemand": "1.738"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.388",
            "yrTerm3.allUpfront": "0.2688",
            "yrTerm1.allUpfront": "0.3314",
            "yrTerm1.partialUpfront": "0.3383",
            "yrTerm3.partialUpfront": "0.2859"
          },
          "ondemand": "0.474"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.204",
            "yrTerm3.allUpfront": "0.1218",
            "yrTerm1.allUpfront": "0.1726",
            "yrTerm1.partialUpfront": "0.1758",
            "yrTerm3.partialUpfront": "0.13"
          },
          "ondemand": "0.275"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.589",
            "yrTerm3.allUpfront": "0.3489",
            "yrTerm1.allUpfront": "0.4937",
            "yrTerm1.partialUpfront": "0.5037",
            "yrTerm3.partialUpfront": "0.3714"
          },
          "ondemand": "0.776"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.27",
            "yrTerm3.allUpfront": "0.77",
            "yrTerm1.allUpfront": "1.0637",
            "yrTerm1.partialUpfront": "1.0858",
            "yrTerm3.partialUpfront": "0.8194"
          },
          "ondemand": "1.672"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.339",
            "yrTerm3.allUpfront": "0.226",
            "yrTerm1.allUpfront": "0.2887",
            "yrTerm1.partialUpfront": "0.2944",
            "yrTerm3.partialUpfront": "0.2404"
          },
          "ondemand": "0.386"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.155",
            "yrTerm3.allUpfront": "0.0816",
            "yrTerm1.allUpfront": "0.1314",
            "yrTerm1.partialUpfront": "0.1344",
            "yrTerm3.partialUpfront": "0.0872"
          },
          "ondemand": "0.209"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.633",
            "yrTerm3.allUpfront": "0.3811",
            "yrTerm1.allUpfront": "0.5298",
            "yrTerm1.partialUpfront": "0.5411",
            "yrTerm3.partialUpfront": "0.4056"
          },
          "ondemand": "0.829"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.314",
            "yrTerm3.allUpfront": "0.8022",
            "yrTerm1.allUpfront": "1.1003",
            "yrTerm1.partialUpfront": "1.1233",
            "yrTerm3.partialUpfront": "0.8536"
          },
          "ondemand": "1.725"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.381",
            "yrTerm3.allUpfront": "0.2608",
            "yrTerm1.allUpfront": "0.3271",
            "yrTerm1.partialUpfront": "0.3338",
            "yrTerm3.partialUpfront": "0.2775"
          },
          "ondemand": "0.442"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.196",
            "yrTerm3.allUpfront": "0.1139",
            "yrTerm1.allUpfront": "0.1675",
            "yrTerm1.partialUpfront": "0.1715",
            "yrTerm3.partialUpfront": "0.1214"
          },
          "ondemand": "0.262"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.638",
            "yrTerm3.allUpfront": "0.3916",
            "yrTerm1.allUpfront": "0.5345",
            "yrTerm1.partialUpfront": "0.5446",
            "yrTerm3.partialUpfront": "0.4171"
          },
          "ondemand": "0.834"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.319",
            "yrTerm3.allUpfront": "0.8809",
            "yrTerm1.allUpfront": "1.1045",
            "yrTerm1.partialUpfront": "1.1267",
            "yrTerm3.partialUpfront": "0.9376"
          },
          "ondemand": "1.73"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.386",
            "yrTerm3.allUpfront": "0.271",
            "yrTerm1.allUpfront": "0.3314",
            "yrTerm1.partialUpfront": "0.3383",
            "yrTerm3.partialUpfront": "0.2883"
          },
          "ondemand": "0.447"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.2",
            "yrTerm3.allUpfront": "0.1242",
            "yrTerm1.allUpfront": "0.1725",
            "yrTerm1.partialUpfront": "0.1757",
            "yrTerm3.partialUpfront": "0.1326"
          },
          "ondemand": "0.267"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.615",
            "yrTerm3.allUpfront": "0.3745",
            "yrTerm1.allUpfront": "0.5146",
            "yrTerm1.partialUpfront": "0.5254",
            "yrTerm3.partialUpfront": "0.3984"
          },
          "ondemand": "0.805"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.296",
            "yrTerm3.allUpfront": "0.7956",
            "yrTerm1.allUpfront": "1.0846",
            "yrTerm1.partialUpfront": "1.1075",
            "yrTerm3.partialUpfront": "0.8464"
          },
          "ondemand": "1.701"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.363",
            "yrTerm3.allUpfront": "0.254",
            "yrTerm1.allUpfront": "0.3106",
            "yrTerm1.partialUpfront": "0.3177",
            "yrTerm3.partialUpfront": "0.2702"
          },
          "ondemand": "0.417"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.179",
            "yrTerm3.allUpfront": "0.108",
            "yrTerm1.allUpfront": "0.1523",
            "yrTerm1.partialUpfront": "0.1556",
            "yrTerm3.partialUpfront": "0.1148"
          },
          "ondemand": "0.238"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 15,
      "max_enis": 4
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": true,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": true,
    "storage": null,
    "instance_type": "c4.xlarge",
    "ECU": 16.0,
    "memory": 7.5,
    "ebs_max_bandwidth": 750.0
  },
  {
    "family": "Compute optimized",
    "enhanced_networking": true,
    "vCPU": 8,
    "generation": "current",
    "ebs_iops": 125.0,
    "network_performance": "High",
    "ebs_throughput": 8000.0,
    "pretty_name": "C4 High-CPU Double Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.226",
            "yrTerm3.allUpfront": "0.7264",
            "yrTerm1.allUpfront": "1.0256",
            "yrTerm1.partialUpfront": "1.0475",
            "yrTerm3.partialUpfront": "0.7724"
          },
          "ondemand": "1.616"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.525",
            "yrTerm3.allUpfront": "1.5292",
            "yrTerm1.allUpfront": "2.1136",
            "yrTerm1.partialUpfront": "2.1575",
            "yrTerm3.partialUpfront": "1.6264"
          },
          "ondemand": "3.324"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.679",
            "yrTerm3.allUpfront": "0.4514",
            "yrTerm1.allUpfront": "0.5795",
            "yrTerm1.partialUpfront": "0.5907",
            "yrTerm3.partialUpfront": "0.4803"
          },
          "ondemand": "0.773"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.311",
            "yrTerm3.allUpfront": "0.1632",
            "yrTerm1.allUpfront": "0.2629",
            "yrTerm1.partialUpfront": "0.2688",
            "yrTerm3.partialUpfront": "0.1733"
          },
          "ondemand": "0.419"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.338",
            "yrTerm3.allUpfront": "0.8191",
            "yrTerm1.allUpfront": "1.1186",
            "yrTerm1.partialUpfront": "1.1415",
            "yrTerm3.partialUpfront": "0.8715"
          },
          "ondemand": "1.728"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.637",
            "yrTerm3.allUpfront": "1.6218",
            "yrTerm1.allUpfront": "2.2066",
            "yrTerm1.partialUpfront": "2.2516",
            "yrTerm3.partialUpfront": "1.7256"
          },
          "ondemand": "3.436"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.788",
            "yrTerm3.allUpfront": "0.5489",
            "yrTerm1.allUpfront": "0.6764",
            "yrTerm1.partialUpfront": "0.6902",
            "yrTerm3.partialUpfront": "0.584"
          },
          "ondemand": "0.929"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.405",
            "yrTerm3.allUpfront": "0.2556",
            "yrTerm1.allUpfront": "0.3553",
            "yrTerm1.partialUpfront": "0.3632",
            "yrTerm3.partialUpfront": "0.2722"
          },
          "ondemand": "0.531"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.8962",
            "yrTerm3.allUpfront": "0.7441",
            "yrTerm1.allUpfront": "0.8342",
            "yrTerm1.partialUpfront": "0.8407",
            "yrTerm3.partialUpfront": "0.7588"
          },
          "ondemand": "0.992"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.1662",
            "yrTerm3.allUpfront": "2.0141",
            "yrTerm1.allUpfront": "2.1042",
            "yrTerm1.partialUpfront": "2.1107",
            "yrTerm3.partialUpfront": "2.0288"
          },
          "ondemand": "2.262"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.7502",
            "yrTerm3.allUpfront": "0.5981",
            "yrTerm1.allUpfront": "0.6882",
            "yrTerm1.partialUpfront": "0.6947",
            "yrTerm3.partialUpfront": "0.6128"
          },
          "ondemand": "0.846"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.3822",
            "yrTerm3.allUpfront": "0.2301",
            "yrTerm1.allUpfront": "0.3202",
            "yrTerm1.partialUpfront": "0.3267",
            "yrTerm3.partialUpfront": "0.2448"
          },
          "ondemand": "0.478"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.345",
            "yrTerm3.allUpfront": "0.8068",
            "yrTerm1.allUpfront": "1.1271",
            "yrTerm1.partialUpfront": "1.1499",
            "yrTerm3.partialUpfront": "0.8581"
          },
          "ondemand": "1.775"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.644",
            "yrTerm3.allUpfront": "1.6095",
            "yrTerm1.allUpfront": "2.215",
            "yrTerm1.partialUpfront": "2.26",
            "yrTerm3.partialUpfront": "1.7121"
          },
          "ondemand": "3.483"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.797",
            "yrTerm3.allUpfront": "0.536",
            "yrTerm1.allUpfront": "0.6857",
            "yrTerm1.partialUpfront": "0.6998",
            "yrTerm3.partialUpfront": "0.5702"
          },
          "ondemand": "0.976"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.429",
            "yrTerm3.allUpfront": "0.2436",
            "yrTerm1.allUpfront": "0.3635",
            "yrTerm1.partialUpfront": "0.3708",
            "yrTerm3.partialUpfront": "0.259"
          },
          "ondemand": "0.578"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.324",
            "yrTerm3.allUpfront": "0.8068",
            "yrTerm1.allUpfront": "1.1087",
            "yrTerm1.partialUpfront": "1.1318",
            "yrTerm3.partialUpfront": "0.8581"
          },
          "ondemand": "1.746"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.623",
            "yrTerm3.allUpfront": "1.6095",
            "yrTerm1.allUpfront": "2.1967",
            "yrTerm1.partialUpfront": "2.2419",
            "yrTerm3.partialUpfront": "1.7121"
          },
          "ondemand": "3.454"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.776",
            "yrTerm3.allUpfront": "0.536",
            "yrTerm1.allUpfront": "0.6669",
            "yrTerm1.partialUpfront": "0.6802",
            "yrTerm3.partialUpfront": "0.5702"
          },
          "ondemand": "0.947"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.408",
            "yrTerm3.allUpfront": "0.2436",
            "yrTerm1.allUpfront": "0.3453",
            "yrTerm1.partialUpfront": "0.3526",
            "yrTerm3.partialUpfront": "0.259"
          },
          "ondemand": "0.549"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.226",
            "yrTerm3.allUpfront": "0.7264",
            "yrTerm1.allUpfront": "1.0256",
            "yrTerm1.partialUpfront": "1.0475",
            "yrTerm3.partialUpfront": "0.7724"
          },
          "ondemand": "1.616"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.525",
            "yrTerm3.allUpfront": "1.5292",
            "yrTerm1.allUpfront": "2.1136",
            "yrTerm1.partialUpfront": "2.1575",
            "yrTerm3.partialUpfront": "1.6264"
          },
          "ondemand": "3.324"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.679",
            "yrTerm3.allUpfront": "0.4514",
            "yrTerm1.allUpfront": "0.5795",
            "yrTerm1.partialUpfront": "0.5907",
            "yrTerm3.partialUpfront": "0.4803"
          },
          "ondemand": "0.773"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.311",
            "yrTerm3.allUpfront": "0.1632",
            "yrTerm1.allUpfront": "0.2629",
            "yrTerm1.partialUpfront": "0.2688",
            "yrTerm3.partialUpfront": "0.1733"
          },
          "ondemand": "0.419"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.311",
            "yrTerm3.allUpfront": "0.7909",
            "yrTerm1.allUpfront": "1.0979",
            "yrTerm1.partialUpfront": "1.1203",
            "yrTerm3.partialUpfront": "0.8418"
          },
          "ondemand": "1.721"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.61",
            "yrTerm3.allUpfront": "1.5936",
            "yrTerm1.allUpfront": "2.1858",
            "yrTerm1.partialUpfront": "2.2304",
            "yrTerm3.partialUpfront": "1.6958"
          },
          "ondemand": "3.429"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.762",
            "yrTerm3.allUpfront": "0.5193",
            "yrTerm1.allUpfront": "0.6561",
            "yrTerm1.partialUpfront": "0.6695",
            "yrTerm3.partialUpfront": "0.5524"
          },
          "ondemand": "0.884"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.391",
            "yrTerm3.allUpfront": "0.2279",
            "yrTerm1.allUpfront": "0.3349",
            "yrTerm1.partialUpfront": "0.3419",
            "yrTerm3.partialUpfront": "0.2429"
          },
          "ondemand": "0.524"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.324",
            "yrTerm3.allUpfront": "0.8118",
            "yrTerm1.allUpfront": "1.1078",
            "yrTerm1.partialUpfront": "1.1302",
            "yrTerm3.partialUpfront": "0.8637"
          },
          "ondemand": "1.731"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.623",
            "yrTerm3.allUpfront": "1.7505",
            "yrTerm1.allUpfront": "2.1958",
            "yrTerm1.partialUpfront": "2.2402",
            "yrTerm3.partialUpfront": "1.8622"
          },
          "ondemand": "3.439"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.775",
            "yrTerm3.allUpfront": "0.5414",
            "yrTerm1.allUpfront": "0.6669",
            "yrTerm1.partialUpfront": "0.6801",
            "yrTerm3.partialUpfront": "0.5759"
          },
          "ondemand": "0.894"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.401",
            "yrTerm3.allUpfront": "0.2484",
            "yrTerm1.allUpfront": "0.345",
            "yrTerm1.partialUpfront": "0.3521",
            "yrTerm3.partialUpfront": "0.2643"
          },
          "ondemand": "0.534"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.274",
            "yrTerm3.allUpfront": "0.7786",
            "yrTerm1.allUpfront": "1.0671",
            "yrTerm1.partialUpfront": "1.0884",
            "yrTerm3.partialUpfront": "0.8283"
          },
          "ondemand": "1.674"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.573",
            "yrTerm3.allUpfront": "1.5813",
            "yrTerm1.allUpfront": "2.1551",
            "yrTerm1.partialUpfront": "2.2",
            "yrTerm3.partialUpfront": "1.6823"
          },
          "ondemand": "3.382"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.726",
            "yrTerm3.allUpfront": "0.5073",
            "yrTerm1.allUpfront": "0.6241",
            "yrTerm1.partialUpfront": "0.6372",
            "yrTerm3.partialUpfront": "0.5397"
          },
          "ondemand": "0.834"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.357",
            "yrTerm3.allUpfront": "0.2159",
            "yrTerm1.allUpfront": "0.3046",
            "yrTerm1.partialUpfront": "0.3104",
            "yrTerm3.partialUpfront": "0.2297"
          },
          "ondemand": "0.477"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 15,
      "max_enis": 4
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": true,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": true,
    "storage": null,
    "instance_type": "c4.2xlarge",
    "ECU": 31.0,
    "memory": 15.0,
    "ebs_max_bandwidth": 1000.0
  },
  {
    "family": "Compute optimized",
    "enhanced_networking": true,
    "vCPU": 16,
    "generation": "current",
    "ebs_iops": 250.0,
    "network_performance": "High",
    "ebs_throughput": 16000.0,
    "pretty_name": "C4 High-CPU Quadruple Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.657",
            "yrTerm3.allUpfront": "1.0171",
            "yrTerm1.allUpfront": "1.3872",
            "yrTerm1.partialUpfront": "1.4159",
            "yrTerm3.partialUpfront": "1.0822"
          },
          "ondemand": "2.189"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "4.203",
            "yrTerm3.allUpfront": "2.5357",
            "yrTerm1.allUpfront": "3.5202",
            "yrTerm1.partialUpfront": "3.5919",
            "yrTerm3.partialUpfront": "2.6978"
          },
          "ondemand": "5.538"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.357",
            "yrTerm3.allUpfront": "0.9027",
            "yrTerm1.allUpfront": "1.1598",
            "yrTerm1.partialUpfront": "1.184",
            "yrTerm3.partialUpfront": "0.9604"
          },
          "ondemand": "1.546"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.621",
            "yrTerm3.allUpfront": "0.3265",
            "yrTerm1.allUpfront": "0.5259",
            "yrTerm1.partialUpfront": "0.5365",
            "yrTerm3.partialUpfront": "0.3476"
          },
          "ondemand": "0.838"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.88",
            "yrTerm3.allUpfront": "1.2024",
            "yrTerm1.allUpfront": "1.5721",
            "yrTerm1.partialUpfront": "1.6038",
            "yrTerm3.partialUpfront": "1.2796"
          },
          "ondemand": "2.412"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "4.426",
            "yrTerm3.allUpfront": "2.7201",
            "yrTerm1.allUpfront": "3.7042",
            "yrTerm1.partialUpfront": "3.78",
            "yrTerm3.partialUpfront": "2.8941"
          },
          "ondemand": "5.761"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.576",
            "yrTerm3.allUpfront": "1.0978",
            "yrTerm1.allUpfront": "1.3547",
            "yrTerm1.partialUpfront": "1.3821",
            "yrTerm3.partialUpfront": "1.1678"
          },
          "ondemand": "1.858"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.808",
            "yrTerm3.allUpfront": "0.5113",
            "yrTerm1.allUpfront": "0.7104",
            "yrTerm1.partialUpfront": "0.7253",
            "yrTerm3.partialUpfront": "0.5444"
          },
          "ondemand": "1.061"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.7914",
            "yrTerm3.allUpfront": "1.4882",
            "yrTerm1.allUpfront": "1.6675",
            "yrTerm1.partialUpfront": "1.6805",
            "yrTerm3.partialUpfront": "1.5176"
          },
          "ondemand": "1.983"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "4.3314",
            "yrTerm3.allUpfront": "4.0282",
            "yrTerm1.allUpfront": "4.2075",
            "yrTerm1.partialUpfront": "4.2205",
            "yrTerm3.partialUpfront": "4.0576"
          },
          "ondemand": "4.523"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.4994",
            "yrTerm3.allUpfront": "1.1962",
            "yrTerm1.allUpfront": "1.3755",
            "yrTerm1.partialUpfront": "1.3885",
            "yrTerm3.partialUpfront": "1.2256"
          },
          "ondemand": "1.691"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.7634",
            "yrTerm3.allUpfront": "0.4602",
            "yrTerm1.allUpfront": "0.6395",
            "yrTerm1.partialUpfront": "0.6525",
            "yrTerm3.partialUpfront": "0.4896"
          },
          "ondemand": "0.955"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.897",
            "yrTerm3.allUpfront": "1.1778",
            "yrTerm1.allUpfront": "1.5885",
            "yrTerm1.partialUpfront": "1.6205",
            "yrTerm3.partialUpfront": "1.2526"
          },
          "ondemand": "2.506"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "4.443",
            "yrTerm3.allUpfront": "2.6964",
            "yrTerm1.allUpfront": "3.7211",
            "yrTerm1.partialUpfront": "3.7966",
            "yrTerm3.partialUpfront": "2.8681"
          },
          "ondemand": "5.855"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.595",
            "yrTerm3.allUpfront": "1.0719",
            "yrTerm1.allUpfront": "1.372",
            "yrTerm1.partialUpfront": "1.4",
            "yrTerm3.partialUpfront": "1.1403"
          },
          "ondemand": "1.952"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.859",
            "yrTerm3.allUpfront": "0.4873",
            "yrTerm1.allUpfront": "0.7271",
            "yrTerm1.partialUpfront": "0.7416",
            "yrTerm3.partialUpfront": "0.518"
          },
          "ondemand": "1.155"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.854",
            "yrTerm3.allUpfront": "1.1778",
            "yrTerm1.allUpfront": "1.5524",
            "yrTerm1.partialUpfront": "1.5839",
            "yrTerm3.partialUpfront": "1.2526"
          },
          "ondemand": "2.448"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "4.4",
            "yrTerm3.allUpfront": "2.6964",
            "yrTerm1.allUpfront": "3.6845",
            "yrTerm1.partialUpfront": "3.759",
            "yrTerm3.partialUpfront": "2.8681"
          },
          "ondemand": "5.797"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.552",
            "yrTerm3.allUpfront": "1.0719",
            "yrTerm1.allUpfront": "1.3338",
            "yrTerm1.partialUpfront": "1.3605",
            "yrTerm3.partialUpfront": "1.1403"
          },
          "ondemand": "1.894"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.816",
            "yrTerm3.allUpfront": "0.4873",
            "yrTerm1.allUpfront": "0.6908",
            "yrTerm1.partialUpfront": "0.7041",
            "yrTerm3.partialUpfront": "0.518"
          },
          "ondemand": "1.097"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.657",
            "yrTerm3.allUpfront": "1.0171",
            "yrTerm1.allUpfront": "1.3872",
            "yrTerm1.partialUpfront": "1.4159",
            "yrTerm3.partialUpfront": "1.0822"
          },
          "ondemand": "2.189"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "4.203",
            "yrTerm3.allUpfront": "2.5357",
            "yrTerm1.allUpfront": "3.5202",
            "yrTerm1.partialUpfront": "3.5919",
            "yrTerm3.partialUpfront": "2.6978"
          },
          "ondemand": "5.538"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.357",
            "yrTerm3.allUpfront": "0.9027",
            "yrTerm1.allUpfront": "1.1598",
            "yrTerm1.partialUpfront": "1.184",
            "yrTerm3.partialUpfront": "0.9604"
          },
          "ondemand": "1.546"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.621",
            "yrTerm3.allUpfront": "0.3265",
            "yrTerm1.allUpfront": "0.5259",
            "yrTerm1.partialUpfront": "0.5365",
            "yrTerm3.partialUpfront": "0.3476"
          },
          "ondemand": "0.838"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.829",
            "yrTerm3.allUpfront": "1.147",
            "yrTerm1.allUpfront": "1.5311",
            "yrTerm1.partialUpfront": "1.5623",
            "yrTerm3.partialUpfront": "1.22"
          },
          "ondemand": "2.4"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "4.375",
            "yrTerm3.allUpfront": "2.6646",
            "yrTerm1.allUpfront": "3.6637",
            "yrTerm1.partialUpfront": "3.7384",
            "yrTerm3.partialUpfront": "2.8345"
          },
          "ondemand": "5.749"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.525",
            "yrTerm3.allUpfront": "1.0394",
            "yrTerm1.allUpfront": "1.3115",
            "yrTerm1.partialUpfront": "1.3383",
            "yrTerm3.partialUpfront": "1.1057"
          },
          "ondemand": "1.768"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.783",
            "yrTerm3.allUpfront": "0.4558",
            "yrTerm1.allUpfront": "0.6697",
            "yrTerm1.partialUpfront": "0.6838",
            "yrTerm3.partialUpfront": "0.4848"
          },
          "ondemand": "1.049"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.854",
            "yrTerm3.allUpfront": "1.188",
            "yrTerm1.allUpfront": "1.5505",
            "yrTerm1.partialUpfront": "1.5813",
            "yrTerm3.partialUpfront": "1.2638"
          },
          "ondemand": "2.42"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "4.402",
            "yrTerm3.allUpfront": "2.9385",
            "yrTerm1.allUpfront": "3.6845",
            "yrTerm1.partialUpfront": "3.7599",
            "yrTerm3.partialUpfront": "3.1261"
          },
          "ondemand": "5.769"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.55",
            "yrTerm3.allUpfront": "1.0826",
            "yrTerm1.allUpfront": "1.3328",
            "yrTerm1.partialUpfront": "1.3603",
            "yrTerm3.partialUpfront": "1.1517"
          },
          "ondemand": "1.789"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.802",
            "yrTerm3.allUpfront": "0.4968",
            "yrTerm1.allUpfront": "0.6898",
            "yrTerm1.partialUpfront": "0.7033",
            "yrTerm3.partialUpfront": "0.5286"
          },
          "ondemand": "1.069"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.756",
            "yrTerm3.allUpfront": "1.1225",
            "yrTerm1.allUpfront": "1.4702",
            "yrTerm1.partialUpfront": "1.5012",
            "yrTerm3.partialUpfront": "1.1941"
          },
          "ondemand": "2.304"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "4.302",
            "yrTerm3.allUpfront": "2.641",
            "yrTerm1.allUpfront": "3.6032",
            "yrTerm1.partialUpfront": "3.6774",
            "yrTerm3.partialUpfront": "2.8096"
          },
          "ondemand": "5.653"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.452",
            "yrTerm3.allUpfront": "1.0136",
            "yrTerm1.allUpfront": "1.2473",
            "yrTerm1.partialUpfront": "1.2735",
            "yrTerm3.partialUpfront": "1.0783"
          },
          "ondemand": "1.667"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.714",
            "yrTerm3.allUpfront": "0.4318",
            "yrTerm1.allUpfront": "0.609",
            "yrTerm1.partialUpfront": "0.6217",
            "yrTerm3.partialUpfront": "0.4594"
          },
          "ondemand": "0.953"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 30,
      "max_enis": 8
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": true,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": true,
    "storage": null,
    "instance_type": "c4.4xlarge",
    "ECU": 62.0,
    "memory": 30.0,
    "ebs_max_bandwidth": 2000.0
  },
  {
    "family": "Compute optimized",
    "enhanced_networking": true,
    "vCPU": 36,
    "generation": "current",
    "ebs_iops": 500.0,
    "network_performance": "10 Gigabit",
    "ebs_throughput": 32000.0,
    "pretty_name": "C4 High-CPU Eight Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.154",
            "yrTerm3.allUpfront": "2.0352",
            "yrTerm1.allUpfront": "2.6404",
            "yrTerm1.partialUpfront": "2.6939",
            "yrTerm3.partialUpfront": "2.1655"
          },
          "ondemand": "4.185"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "8.407",
            "yrTerm3.allUpfront": "5.7337",
            "yrTerm1.allUpfront": "7.0404",
            "yrTerm1.partialUpfront": "7.1838",
            "yrTerm3.partialUpfront": "6.1"
          },
          "ondemand": "12.485"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.769",
            "yrTerm3.allUpfront": "1.8049",
            "yrTerm1.allUpfront": "2.3187",
            "yrTerm1.partialUpfront": "2.3662",
            "yrTerm3.partialUpfront": "1.9201"
          },
          "ondemand": "3.091"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.242",
            "yrTerm3.allUpfront": "0.653",
            "yrTerm1.allUpfront": "1.0517",
            "yrTerm1.partialUpfront": "1.073",
            "yrTerm3.partialUpfront": "0.6953"
          },
          "ondemand": "1.675"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.602",
            "yrTerm3.allUpfront": "2.4048",
            "yrTerm1.allUpfront": "3.0095",
            "yrTerm1.partialUpfront": "3.0715",
            "yrTerm3.partialUpfront": "2.5581"
          },
          "ondemand": "4.632"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "8.855",
            "yrTerm3.allUpfront": "6.1033",
            "yrTerm1.allUpfront": "7.4095",
            "yrTerm1.partialUpfront": "7.5613",
            "yrTerm3.partialUpfront": "6.4927"
          },
          "ondemand": "12.932"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "3.237",
            "yrTerm3.allUpfront": "2.1939",
            "yrTerm1.allUpfront": "2.7083",
            "yrTerm1.partialUpfront": "2.7637",
            "yrTerm3.partialUpfront": "2.334"
          },
          "ondemand": "3.778"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.617",
            "yrTerm3.allUpfront": "1.0226",
            "yrTerm1.allUpfront": "1.4209",
            "yrTerm1.partialUpfront": "1.4505",
            "yrTerm3.partialUpfront": "1.0878"
          },
          "ondemand": "2.122"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.8409",
            "yrTerm3.allUpfront": "3.2334",
            "yrTerm1.allUpfront": "3.5928",
            "yrTerm1.partialUpfront": "3.6189",
            "yrTerm3.partialUpfront": "3.2922"
          },
          "ondemand": "4.223"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "9.5559",
            "yrTerm3.allUpfront": "8.9484",
            "yrTerm1.allUpfront": "9.3078",
            "yrTerm1.partialUpfront": "9.3339",
            "yrTerm3.partialUpfront": "9.0072"
          },
          "ondemand": "9.938"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "3.1839",
            "yrTerm3.allUpfront": "2.5764",
            "yrTerm1.allUpfront": "2.9358",
            "yrTerm1.partialUpfront": "2.9619",
            "yrTerm3.partialUpfront": "2.6352"
          },
          "ondemand": "3.566"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.5279",
            "yrTerm3.allUpfront": "0.9204",
            "yrTerm1.allUpfront": "1.2798",
            "yrTerm1.partialUpfront": "1.3059",
            "yrTerm3.partialUpfront": "0.9792"
          },
          "ondemand": "1.91"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.633",
            "yrTerm3.allUpfront": "2.3566",
            "yrTerm1.allUpfront": "3.0424",
            "yrTerm1.partialUpfront": "3.1045",
            "yrTerm3.partialUpfront": "2.5073"
          },
          "ondemand": "4.82"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "8.886",
            "yrTerm3.allUpfront": "6.0551",
            "yrTerm1.allUpfront": "7.4421",
            "yrTerm1.partialUpfront": "7.5941",
            "yrTerm3.partialUpfront": "6.4418"
          },
          "ondemand": "13.12"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "3.274",
            "yrTerm3.allUpfront": "2.1433",
            "yrTerm1.allUpfront": "2.7424",
            "yrTerm1.partialUpfront": "2.7983",
            "yrTerm3.partialUpfront": "2.2801"
          },
          "ondemand": "3.966"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.718",
            "yrTerm3.allUpfront": "0.9746",
            "yrTerm1.allUpfront": "1.4541",
            "yrTerm1.partialUpfront": "1.4842",
            "yrTerm3.partialUpfront": "1.037"
          },
          "ondemand": "2.31"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.547",
            "yrTerm3.allUpfront": "2.3566",
            "yrTerm1.allUpfront": "2.9697",
            "yrTerm1.partialUpfront": "3.0302",
            "yrTerm3.partialUpfront": "2.5073"
          },
          "ondemand": "4.705"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "8.8",
            "yrTerm3.allUpfront": "6.0551",
            "yrTerm1.allUpfront": "7.3697",
            "yrTerm1.partialUpfront": "7.52",
            "yrTerm3.partialUpfront": "6.4418"
          },
          "ondemand": "13.005"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "3.184",
            "yrTerm3.allUpfront": "2.1433",
            "yrTerm1.allUpfront": "2.6655",
            "yrTerm1.partialUpfront": "2.72",
            "yrTerm3.partialUpfront": "2.2801"
          },
          "ondemand": "3.851"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.632",
            "yrTerm3.allUpfront": "0.9746",
            "yrTerm1.allUpfront": "1.3814",
            "yrTerm1.partialUpfront": "1.4102",
            "yrTerm3.partialUpfront": "1.037"
          },
          "ondemand": "2.195"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.154",
            "yrTerm3.allUpfront": "2.0352",
            "yrTerm1.allUpfront": "2.6404",
            "yrTerm1.partialUpfront": "2.6939",
            "yrTerm3.partialUpfront": "2.1655"
          },
          "ondemand": "4.185"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "8.407",
            "yrTerm3.allUpfront": "5.7337",
            "yrTerm1.allUpfront": "7.0404",
            "yrTerm1.partialUpfront": "7.1838",
            "yrTerm3.partialUpfront": "6.1"
          },
          "ondemand": "12.485"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.769",
            "yrTerm3.allUpfront": "1.8049",
            "yrTerm1.allUpfront": "2.3187",
            "yrTerm1.partialUpfront": "2.3662",
            "yrTerm3.partialUpfront": "1.9201"
          },
          "ondemand": "3.091"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.242",
            "yrTerm3.allUpfront": "0.653",
            "yrTerm1.allUpfront": "1.0517",
            "yrTerm1.partialUpfront": "1.073",
            "yrTerm3.partialUpfront": "0.6953"
          },
          "ondemand": "1.675"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.498",
            "yrTerm3.allUpfront": "2.293",
            "yrTerm1.allUpfront": "2.9284",
            "yrTerm1.partialUpfront": "2.9882",
            "yrTerm3.partialUpfront": "2.439"
          },
          "ondemand": "4.608"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "8.751",
            "yrTerm3.allUpfront": "5.9924",
            "yrTerm1.allUpfront": "7.3282",
            "yrTerm1.partialUpfront": "7.4778",
            "yrTerm3.partialUpfront": "6.3745"
          },
          "ondemand": "12.908"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "3.131",
            "yrTerm3.allUpfront": "2.0773",
            "yrTerm1.allUpfront": "2.6226",
            "yrTerm1.partialUpfront": "2.6761",
            "yrTerm3.partialUpfront": "2.2099"
          },
          "ondemand": "3.536"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.566",
            "yrTerm3.allUpfront": "0.9115",
            "yrTerm1.allUpfront": "1.3395",
            "yrTerm1.partialUpfront": "1.3677",
            "yrTerm3.partialUpfront": "0.9695"
          },
          "ondemand": "2.098"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.548",
            "yrTerm3.allUpfront": "2.3752",
            "yrTerm1.allUpfront": "2.9679",
            "yrTerm1.partialUpfront": "3.0288",
            "yrTerm3.partialUpfront": "2.5267"
          },
          "ondemand": "4.648"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "8.801",
            "yrTerm3.allUpfront": "6.0745",
            "yrTerm1.allUpfront": "7.3679",
            "yrTerm1.partialUpfront": "7.5182",
            "yrTerm3.partialUpfront": "6.4623"
          },
          "ondemand": "12.948"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "3.183",
            "yrTerm3.allUpfront": "2.1637",
            "yrTerm1.allUpfront": "2.6646",
            "yrTerm1.partialUpfront": "2.7181",
            "yrTerm3.partialUpfront": "2.3018"
          },
          "ondemand": "3.578"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.603",
            "yrTerm3.allUpfront": "0.9937",
            "yrTerm1.allUpfront": "1.3797",
            "yrTerm1.partialUpfront": "1.4086",
            "yrTerm3.partialUpfront": "1.0572"
          },
          "ondemand": "2.138"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.353",
            "yrTerm3.allUpfront": "2.2449",
            "yrTerm1.allUpfront": "2.8064",
            "yrTerm1.partialUpfront": "2.8637",
            "yrTerm3.partialUpfront": "2.3881"
          },
          "ondemand": "4.416"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "8.605",
            "yrTerm3.allUpfront": "5.9443",
            "yrTerm1.allUpfront": "7.2054",
            "yrTerm1.partialUpfront": "7.351",
            "yrTerm3.partialUpfront": "6.3237"
          },
          "ondemand": "12.716"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.979",
            "yrTerm3.allUpfront": "2.0266",
            "yrTerm1.allUpfront": "2.4935",
            "yrTerm1.partialUpfront": "2.5454",
            "yrTerm3.partialUpfront": "2.1559"
          },
          "ondemand": "3.334"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.43",
            "yrTerm3.allUpfront": "0.8635",
            "yrTerm1.allUpfront": "1.2182",
            "yrTerm1.partialUpfront": "1.2425",
            "yrTerm3.partialUpfront": "0.9187"
          },
          "ondemand": "1.906"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 30,
      "max_enis": 8
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": true,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": true,
    "storage": null,
    "instance_type": "c4.8xlarge",
    "ECU": 132.0,
    "memory": 60.0,
    "ebs_max_bandwidth": 4000.0
  },
  {
    "family": "Compute optimized",
    "enhanced_networking": true,
    "vCPU": 2,
    "generation": "current",
    "ebs_iops": 0,
    "network_performance": "Moderate",
    "ebs_throughput": 0,
    "pretty_name": "C3 High-CPU Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.192",
            "yrTerm3.allUpfront": "0.1255",
            "yrTerm1.allUpfront": "0.1605",
            "yrTerm1.partialUpfront": "0.1638",
            "yrTerm3.partialUpfront": "0.1335"
          },
          "ondemand": "0.271"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.406",
            "yrTerm3.allUpfront": "0.2944",
            "yrTerm1.allUpfront": "0.3398",
            "yrTerm1.partialUpfront": "0.3468",
            "yrTerm3.partialUpfront": "0.3132"
          },
          "ondemand": "0.561"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.165",
            "yrTerm3.allUpfront": "0.1083",
            "yrTerm1.allUpfront": "0.1379",
            "yrTerm1.partialUpfront": "0.1407",
            "yrTerm3.partialUpfront": "0.1152"
          },
          "ondemand": "0.188"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.073",
            "yrTerm3.allUpfront": "0.0388",
            "yrTerm1.allUpfront": "0.0619",
            "yrTerm1.partialUpfront": "0.0632",
            "yrTerm3.partialUpfront": "0.0413"
          },
          "ondemand": "0.105"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.224",
            "yrTerm3.allUpfront": "0.1495",
            "yrTerm1.allUpfront": "0.188",
            "yrTerm1.partialUpfront": "0.1918",
            "yrTerm3.partialUpfront": "0.159"
          },
          "ondemand": "0.271"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.459",
            "yrTerm3.allUpfront": "0.3357",
            "yrTerm1.allUpfront": "0.3844",
            "yrTerm1.partialUpfront": "0.3922",
            "yrTerm3.partialUpfront": "0.3571"
          },
          "ondemand": "0.572"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.187",
            "yrTerm3.allUpfront": "0.1275",
            "yrTerm1.allUpfront": "0.1563",
            "yrTerm1.partialUpfront": "0.1595",
            "yrTerm3.partialUpfront": "0.1356"
          },
          "ondemand": "0.231"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.097",
            "yrTerm3.allUpfront": "0.0581",
            "yrTerm1.allUpfront": "0.086",
            "yrTerm1.partialUpfront": "0.0878",
            "yrTerm3.partialUpfront": "0.0619"
          },
          "ondemand": "0.128"
        }
      },
      "sa-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.29",
            "yrTerm3.allUpfront": "0.1565",
            "yrTerm1.allUpfront": "0.2427",
            "yrTerm1.partialUpfront": "0.2476",
            "yrTerm3.partialUpfront": "0.1664"
          },
          "ondemand": "0.329"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.229",
            "yrTerm3.allUpfront": "0.1395",
            "yrTerm1.allUpfront": "0.2163",
            "yrTerm1.partialUpfront": "0.2208",
            "yrTerm3.partialUpfront": "0.1483"
          },
          "ondemand": "0.246"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.123",
            "yrTerm3.allUpfront": "0.0697",
            "yrTerm1.allUpfront": "0.1081",
            "yrTerm1.partialUpfront": "0.1104",
            "yrTerm3.partialUpfront": "0.0742"
          },
          "ondemand": "0.163"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.217",
            "yrTerm3.allUpfront": "0.1438",
            "yrTerm1.allUpfront": "0.1817",
            "yrTerm1.partialUpfront": "0.1854",
            "yrTerm3.partialUpfront": "0.153"
          },
          "ondemand": "0.271"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.433",
            "yrTerm3.allUpfront": "0.3135",
            "yrTerm1.allUpfront": "0.3626",
            "yrTerm1.partialUpfront": "0.3699",
            "yrTerm3.partialUpfront": "0.3335"
          },
          "ondemand": "0.561"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.183",
            "yrTerm3.allUpfront": "0.1215",
            "yrTerm1.allUpfront": "0.1535",
            "yrTerm1.partialUpfront": "0.1567",
            "yrTerm3.partialUpfront": "0.1292"
          },
          "ondemand": "0.238"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.098",
            "yrTerm3.allUpfront": "0.0559",
            "yrTerm1.allUpfront": "0.0829",
            "yrTerm1.partialUpfront": "0.0846",
            "yrTerm3.partialUpfront": "0.0595"
          },
          "ondemand": "0.132"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.217",
            "yrTerm3.allUpfront": "0.1438",
            "yrTerm1.allUpfront": "0.1817",
            "yrTerm1.partialUpfront": "0.1854",
            "yrTerm3.partialUpfront": "0.153"
          },
          "ondemand": "0.271"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.433",
            "yrTerm3.allUpfront": "0.3135",
            "yrTerm1.allUpfront": "0.3626",
            "yrTerm1.partialUpfront": "0.3699",
            "yrTerm3.partialUpfront": "0.3335"
          },
          "ondemand": "0.561"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.183",
            "yrTerm3.allUpfront": "0.1215",
            "yrTerm1.allUpfront": "0.1535",
            "yrTerm1.partialUpfront": "0.1567",
            "yrTerm3.partialUpfront": "0.1292"
          },
          "ondemand": "0.238"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.098",
            "yrTerm3.allUpfront": "0.0559",
            "yrTerm1.allUpfront": "0.0829",
            "yrTerm1.partialUpfront": "0.0846",
            "yrTerm3.partialUpfront": "0.0595"
          },
          "ondemand": "0.132"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.192",
            "yrTerm3.allUpfront": "0.1255",
            "yrTerm1.allUpfront": "0.1605",
            "yrTerm1.partialUpfront": "0.1638",
            "yrTerm3.partialUpfront": "0.1335"
          },
          "ondemand": "0.271"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.406",
            "yrTerm3.allUpfront": "0.2944",
            "yrTerm1.allUpfront": "0.3398",
            "yrTerm1.partialUpfront": "0.3468",
            "yrTerm3.partialUpfront": "0.3132"
          },
          "ondemand": "0.561"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.165",
            "yrTerm3.allUpfront": "0.1083",
            "yrTerm1.allUpfront": "0.1379",
            "yrTerm1.partialUpfront": "0.1407",
            "yrTerm3.partialUpfront": "0.1152"
          },
          "ondemand": "0.188"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.073",
            "yrTerm3.allUpfront": "0.0388",
            "yrTerm1.allUpfront": "0.0619",
            "yrTerm1.partialUpfront": "0.0632",
            "yrTerm3.partialUpfront": "0.0413"
          },
          "ondemand": "0.105"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.207",
            "yrTerm3.allUpfront": "0.1334",
            "yrTerm1.allUpfront": "0.1732",
            "yrTerm1.partialUpfront": "0.1768",
            "yrTerm3.partialUpfront": "0.142"
          },
          "ondemand": "0.292"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.18",
            "yrTerm3.allUpfront": "0.1166",
            "yrTerm1.allUpfront": "0.151",
            "yrTerm1.partialUpfront": "0.1541",
            "yrTerm3.partialUpfront": "0.1241"
          },
          "ondemand": "0.209"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.088",
            "yrTerm3.allUpfront": "0.0466",
            "yrTerm1.allUpfront": "0.0743",
            "yrTerm1.partialUpfront": "0.0756",
            "yrTerm3.partialUpfront": "0.0492"
          },
          "ondemand": "0.126"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.192",
            "yrTerm3.allUpfront": "0.1255",
            "yrTerm1.allUpfront": "0.1605",
            "yrTerm1.partialUpfront": "0.1638",
            "yrTerm3.partialUpfront": "0.1335"
          },
          "ondemand": "0.271"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.406",
            "yrTerm3.allUpfront": "0.2944",
            "yrTerm1.allUpfront": "0.3398",
            "yrTerm1.partialUpfront": "0.3468",
            "yrTerm3.partialUpfront": "0.3132"
          },
          "ondemand": "0.561"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.186",
            "yrTerm3.allUpfront": "0.1233",
            "yrTerm1.allUpfront": "0.1555",
            "yrTerm1.partialUpfront": "0.1587",
            "yrTerm3.partialUpfront": "0.1312"
          },
          "ondemand": "0.188"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.09",
            "yrTerm3.allUpfront": "0.052",
            "yrTerm1.allUpfront": "0.0766",
            "yrTerm1.partialUpfront": "0.0782",
            "yrTerm3.partialUpfront": "0.0553"
          },
          "ondemand": "0.12"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.217",
            "yrTerm3.allUpfront": "0.1424",
            "yrTerm1.allUpfront": "0.1815",
            "yrTerm1.partialUpfront": "0.1852",
            "yrTerm3.partialUpfront": "0.1515"
          },
          "ondemand": "0.437"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.19",
            "yrTerm3.allUpfront": "0.1249",
            "yrTerm1.allUpfront": "0.1588",
            "yrTerm1.partialUpfront": "0.1621",
            "yrTerm3.partialUpfront": "0.1329"
          },
          "ondemand": "0.212"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.096",
            "yrTerm3.allUpfront": "0.0564",
            "yrTerm1.allUpfront": "0.0826",
            "yrTerm1.partialUpfront": "0.0844",
            "yrTerm3.partialUpfront": "0.0601"
          },
          "ondemand": "0.129"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.192",
            "yrTerm3.allUpfront": "0.1255",
            "yrTerm1.allUpfront": "0.1605",
            "yrTerm1.partialUpfront": "0.1638",
            "yrTerm3.partialUpfront": "0.1335"
          },
          "ondemand": "0.271"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.406",
            "yrTerm3.allUpfront": "0.2944",
            "yrTerm1.allUpfront": "0.3398",
            "yrTerm1.partialUpfront": "0.3468",
            "yrTerm3.partialUpfront": "0.3132"
          },
          "ondemand": "0.561"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.186",
            "yrTerm3.allUpfront": "0.1233",
            "yrTerm1.allUpfront": "0.1555",
            "yrTerm1.partialUpfront": "0.1587",
            "yrTerm3.partialUpfront": "0.1312"
          },
          "ondemand": "0.188"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.09",
            "yrTerm3.allUpfront": "0.052",
            "yrTerm1.allUpfront": "0.0766",
            "yrTerm1.partialUpfront": "0.0782",
            "yrTerm3.partialUpfront": "0.0553"
          },
          "ondemand": "0.12"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 10,
      "max_enis": 3
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "HVM",
      "PV"
    ],
    "ebs_optimized": false,
    "storage": {
      "ssd": true,
      "devices": 2,
      "size": 16.0
    },
    "instance_type": "c3.large",
    "ECU": 7.0,
    "memory": 3.75,
    "ebs_max_bandwidth": 0
  },
  {
    "family": "Compute optimized",
    "enhanced_networking": true,
    "vCPU": 4,
    "generation": "current",
    "ebs_iops": 62.5,
    "network_performance": "Moderate",
    "ebs_throughput": 4000.0,
    "pretty_name": "C3 High-CPU Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.386",
            "yrTerm3.allUpfront": "0.2509",
            "yrTerm1.allUpfront": "0.3229",
            "yrTerm1.partialUpfront": "0.3295",
            "yrTerm3.partialUpfront": "0.2669"
          },
          "ondemand": "0.542"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.771",
            "yrTerm3.allUpfront": "0.5593",
            "yrTerm1.allUpfront": "0.6455",
            "yrTerm1.partialUpfront": "0.6587",
            "yrTerm3.partialUpfront": "0.595"
          },
          "ondemand": "1.066"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.329",
            "yrTerm3.allUpfront": "0.2191",
            "yrTerm1.allUpfront": "0.2755",
            "yrTerm1.partialUpfront": "0.2811",
            "yrTerm3.partialUpfront": "0.2331"
          },
          "ondemand": "0.376"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.146",
            "yrTerm3.allUpfront": "0.0786",
            "yrTerm1.allUpfront": "0.1248",
            "yrTerm1.partialUpfront": "0.1274",
            "yrTerm3.partialUpfront": "0.0837"
          },
          "ondemand": "0.21"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.449",
            "yrTerm3.allUpfront": "0.2999",
            "yrTerm1.allUpfront": "0.3761",
            "yrTerm1.partialUpfront": "0.3838",
            "yrTerm3.partialUpfront": "0.319"
          },
          "ondemand": "0.542"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.874",
            "yrTerm3.allUpfront": "0.6391",
            "yrTerm1.allUpfront": "0.7318",
            "yrTerm1.partialUpfront": "0.7467",
            "yrTerm3.partialUpfront": "0.6799"
          },
          "ondemand": "1.141"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.374",
            "yrTerm3.allUpfront": "0.2538",
            "yrTerm1.allUpfront": "0.3134",
            "yrTerm1.partialUpfront": "0.3198",
            "yrTerm3.partialUpfront": "0.27"
          },
          "ondemand": "0.462"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.194",
            "yrTerm3.allUpfront": "0.1154",
            "yrTerm1.allUpfront": "0.1718",
            "yrTerm1.partialUpfront": "0.1754",
            "yrTerm3.partialUpfront": "0.1227"
          },
          "ondemand": "0.255"
        }
      },
      "sa-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.579",
            "yrTerm3.allUpfront": "0.3112",
            "yrTerm1.allUpfront": "0.4854",
            "yrTerm1.partialUpfront": "0.4953",
            "yrTerm3.partialUpfront": "0.331"
          },
          "ondemand": "0.657"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.259",
            "yrTerm3.allUpfront": "0.6752",
            "yrTerm1.allUpfront": "1.0549",
            "yrTerm1.partialUpfront": "1.0765",
            "yrTerm3.partialUpfront": "0.7183"
          },
          "ondemand": "1.294"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.46",
            "yrTerm3.allUpfront": "0.2786",
            "yrTerm1.allUpfront": "0.4355",
            "yrTerm1.partialUpfront": "0.4444",
            "yrTerm3.partialUpfront": "0.2964"
          },
          "ondemand": "0.491"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.246",
            "yrTerm3.allUpfront": "0.1384",
            "yrTerm1.allUpfront": "0.2162",
            "yrTerm1.partialUpfront": "0.2207",
            "yrTerm3.partialUpfront": "0.1473"
          },
          "ondemand": "0.325"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.436",
            "yrTerm3.allUpfront": "0.2868",
            "yrTerm1.allUpfront": "0.3655",
            "yrTerm1.partialUpfront": "0.373",
            "yrTerm3.partialUpfront": "0.3051"
          },
          "ondemand": "0.542"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.824",
            "yrTerm3.allUpfront": "0.5966",
            "yrTerm1.allUpfront": "0.6901",
            "yrTerm1.partialUpfront": "0.7042",
            "yrTerm3.partialUpfront": "0.6347"
          },
          "ondemand": "1.066"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.366",
            "yrTerm3.allUpfront": "0.2435",
            "yrTerm1.allUpfront": "0.3068",
            "yrTerm1.partialUpfront": "0.3131",
            "yrTerm3.partialUpfront": "0.2591"
          },
          "ondemand": "0.477"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.196",
            "yrTerm3.allUpfront": "0.1118",
            "yrTerm1.allUpfront": "0.166",
            "yrTerm1.partialUpfront": "0.1694",
            "yrTerm3.partialUpfront": "0.119"
          },
          "ondemand": "0.265"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.436",
            "yrTerm3.allUpfront": "0.2868",
            "yrTerm1.allUpfront": "0.3655",
            "yrTerm1.partialUpfront": "0.373",
            "yrTerm3.partialUpfront": "0.3051"
          },
          "ondemand": "0.542"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.824",
            "yrTerm3.allUpfront": "0.5966",
            "yrTerm1.allUpfront": "0.6901",
            "yrTerm1.partialUpfront": "0.7042",
            "yrTerm3.partialUpfront": "0.6347"
          },
          "ondemand": "1.066"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.366",
            "yrTerm3.allUpfront": "0.2435",
            "yrTerm1.allUpfront": "0.3068",
            "yrTerm1.partialUpfront": "0.3131",
            "yrTerm3.partialUpfront": "0.2591"
          },
          "ondemand": "0.477"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.196",
            "yrTerm3.allUpfront": "0.1118",
            "yrTerm1.allUpfront": "0.166",
            "yrTerm1.partialUpfront": "0.1694",
            "yrTerm3.partialUpfront": "0.119"
          },
          "ondemand": "0.265"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.386",
            "yrTerm3.allUpfront": "0.2509",
            "yrTerm1.allUpfront": "0.3229",
            "yrTerm1.partialUpfront": "0.3295",
            "yrTerm3.partialUpfront": "0.2669"
          },
          "ondemand": "0.542"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.771",
            "yrTerm3.allUpfront": "0.5593",
            "yrTerm1.allUpfront": "0.6455",
            "yrTerm1.partialUpfront": "0.6587",
            "yrTerm3.partialUpfront": "0.595"
          },
          "ondemand": "1.066"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.329",
            "yrTerm3.allUpfront": "0.2191",
            "yrTerm1.allUpfront": "0.2755",
            "yrTerm1.partialUpfront": "0.2811",
            "yrTerm3.partialUpfront": "0.2331"
          },
          "ondemand": "0.376"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.146",
            "yrTerm3.allUpfront": "0.0786",
            "yrTerm1.allUpfront": "0.1248",
            "yrTerm1.partialUpfront": "0.1274",
            "yrTerm3.partialUpfront": "0.0837"
          },
          "ondemand": "0.21"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.416",
            "yrTerm3.allUpfront": "0.2669",
            "yrTerm1.allUpfront": "0.3485",
            "yrTerm1.partialUpfront": "0.3557",
            "yrTerm3.partialUpfront": "0.2839"
          },
          "ondemand": "0.584"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.801",
            "yrTerm3.allUpfront": "0.5759",
            "yrTerm1.allUpfront": "0.6711",
            "yrTerm1.partialUpfront": "0.6849",
            "yrTerm3.partialUpfront": "0.6127"
          },
          "ondemand": "1.108"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.359",
            "yrTerm3.allUpfront": "0.2349",
            "yrTerm1.allUpfront": "0.3009",
            "yrTerm1.partialUpfront": "0.3071",
            "yrTerm3.partialUpfront": "0.2499"
          },
          "ondemand": "0.418"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.176",
            "yrTerm3.allUpfront": "0.0944",
            "yrTerm1.allUpfront": "0.1498",
            "yrTerm1.partialUpfront": "0.1523",
            "yrTerm3.partialUpfront": "0.1004"
          },
          "ondemand": "0.252"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.386",
            "yrTerm3.allUpfront": "0.2509",
            "yrTerm1.allUpfront": "0.3229",
            "yrTerm1.partialUpfront": "0.3295",
            "yrTerm3.partialUpfront": "0.2669"
          },
          "ondemand": "0.542"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.771",
            "yrTerm3.allUpfront": "0.5593",
            "yrTerm1.allUpfront": "0.6455",
            "yrTerm1.partialUpfront": "0.6587",
            "yrTerm3.partialUpfront": "0.595"
          },
          "ondemand": "1.066"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.371",
            "yrTerm3.allUpfront": "0.2482",
            "yrTerm1.allUpfront": "0.3107",
            "yrTerm1.partialUpfront": "0.3171",
            "yrTerm3.partialUpfront": "0.2641"
          },
          "ondemand": "0.376"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.18",
            "yrTerm3.allUpfront": "0.103",
            "yrTerm1.allUpfront": "0.1542",
            "yrTerm1.partialUpfront": "0.1574",
            "yrTerm3.partialUpfront": "0.1097"
          },
          "ondemand": "0.239"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.436",
            "yrTerm3.allUpfront": "0.2834",
            "yrTerm1.allUpfront": "0.3651",
            "yrTerm1.partialUpfront": "0.3725",
            "yrTerm3.partialUpfront": "0.3015"
          },
          "ondemand": "0.825"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.2",
            "yrTerm3.allUpfront": "0.8666",
            "yrTerm1.allUpfront": "1.0054",
            "yrTerm1.partialUpfront": "1.0258",
            "yrTerm3.partialUpfront": "0.9219"
          },
          "ondemand": "1.721"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.379",
            "yrTerm3.allUpfront": "0.2512",
            "yrTerm1.allUpfront": "0.3176",
            "yrTerm1.partialUpfront": "0.324",
            "yrTerm3.partialUpfront": "0.2673"
          },
          "ondemand": "0.424"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.192",
            "yrTerm3.allUpfront": "0.111",
            "yrTerm1.allUpfront": "0.1666",
            "yrTerm1.partialUpfront": "0.17",
            "yrTerm3.partialUpfront": "0.1181"
          },
          "ondemand": "0.258"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.386",
            "yrTerm3.allUpfront": "0.2509",
            "yrTerm1.allUpfront": "0.3229",
            "yrTerm1.partialUpfront": "0.3295",
            "yrTerm3.partialUpfront": "0.2669"
          },
          "ondemand": "0.542"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.771",
            "yrTerm3.allUpfront": "0.5593",
            "yrTerm1.allUpfront": "0.6455",
            "yrTerm1.partialUpfront": "0.6587",
            "yrTerm3.partialUpfront": "0.595"
          },
          "ondemand": "1.066"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.371",
            "yrTerm3.allUpfront": "0.2482",
            "yrTerm1.allUpfront": "0.3107",
            "yrTerm1.partialUpfront": "0.3171",
            "yrTerm3.partialUpfront": "0.2641"
          },
          "ondemand": "0.376"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.18",
            "yrTerm3.allUpfront": "0.103",
            "yrTerm1.allUpfront": "0.1542",
            "yrTerm1.partialUpfront": "0.1574",
            "yrTerm3.partialUpfront": "0.1097"
          },
          "ondemand": "0.239"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 15,
      "max_enis": 4
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "HVM",
      "PV"
    ],
    "ebs_optimized": true,
    "storage": {
      "ssd": true,
      "devices": 2,
      "size": 40.0
    },
    "instance_type": "c3.xlarge",
    "ECU": 14.0,
    "memory": 7.5,
    "ebs_max_bandwidth": 500.0
  },
  {
    "family": "Compute optimized",
    "enhanced_networking": true,
    "vCPU": 8,
    "generation": "current",
    "ebs_iops": 125.0,
    "network_performance": "High",
    "ebs_throughput": 8000.0,
    "pretty_name": "C3 High-CPU Double Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.771",
            "yrTerm3.allUpfront": "0.5027",
            "yrTerm1.allUpfront": "0.6459",
            "yrTerm1.partialUpfront": "0.6591",
            "yrTerm3.partialUpfront": "0.5348"
          },
          "ondemand": "1.083"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.542",
            "yrTerm3.allUpfront": "1.1176",
            "yrTerm1.allUpfront": "1.2913",
            "yrTerm1.partialUpfront": "1.3176",
            "yrTerm3.partialUpfront": "1.1889"
          },
          "ondemand": "2.131"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.657",
            "yrTerm3.allUpfront": "0.437",
            "yrTerm1.allUpfront": "0.5507",
            "yrTerm1.partialUpfront": "0.5619",
            "yrTerm3.partialUpfront": "0.4649"
          },
          "ondemand": "0.752"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.292",
            "yrTerm3.allUpfront": "0.1572",
            "yrTerm1.allUpfront": "0.2477",
            "yrTerm1.partialUpfront": "0.2529",
            "yrTerm3.partialUpfront": "0.1673"
          },
          "ondemand": "0.42"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.9",
            "yrTerm3.allUpfront": "0.6017",
            "yrTerm1.allUpfront": "0.754",
            "yrTerm1.partialUpfront": "0.7694",
            "yrTerm3.partialUpfront": "0.6401"
          },
          "ondemand": "1.083"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.746",
            "yrTerm3.allUpfront": "1.2772",
            "yrTerm1.allUpfront": "1.4628",
            "yrTerm1.partialUpfront": "1.4926",
            "yrTerm3.partialUpfront": "1.3588"
          },
          "ondemand": "2.282"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.744",
            "yrTerm3.allUpfront": "0.5098",
            "yrTerm1.allUpfront": "0.6232",
            "yrTerm1.partialUpfront": "0.6359",
            "yrTerm3.partialUpfront": "0.5424"
          },
          "ondemand": "0.925"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.388",
            "yrTerm3.allUpfront": "0.2317",
            "yrTerm1.allUpfront": "0.3438",
            "yrTerm1.partialUpfront": "0.3509",
            "yrTerm3.partialUpfront": "0.2465"
          },
          "ondemand": "0.511"
        }
      },
      "sa-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.162",
            "yrTerm3.allUpfront": "0.624",
            "yrTerm1.allUpfront": "0.9737",
            "yrTerm1.partialUpfront": "0.9936",
            "yrTerm3.partialUpfront": "0.6638"
          },
          "ondemand": "1.313"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.519",
            "yrTerm3.allUpfront": "1.3505",
            "yrTerm1.allUpfront": "2.1098",
            "yrTerm1.partialUpfront": "2.1529",
            "yrTerm3.partialUpfront": "1.4366"
          },
          "ondemand": "2.587"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.92",
            "yrTerm3.allUpfront": "0.5573",
            "yrTerm1.allUpfront": "0.8711",
            "yrTerm1.partialUpfront": "0.8889",
            "yrTerm3.partialUpfront": "0.5928"
          },
          "ondemand": "0.982"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.492",
            "yrTerm3.allUpfront": "0.2769",
            "yrTerm1.allUpfront": "0.4324",
            "yrTerm1.partialUpfront": "0.4414",
            "yrTerm3.partialUpfront": "0.2946"
          },
          "ondemand": "0.65"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.87",
            "yrTerm3.allUpfront": "0.5736",
            "yrTerm1.allUpfront": "0.729",
            "yrTerm1.partialUpfront": "0.7438",
            "yrTerm3.partialUpfront": "0.6102"
          },
          "ondemand": "1.083"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.647",
            "yrTerm3.allUpfront": "1.1922",
            "yrTerm1.allUpfront": "1.3792",
            "yrTerm1.partialUpfront": "1.4073",
            "yrTerm3.partialUpfront": "1.2683"
          },
          "ondemand": "2.131"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.731",
            "yrTerm3.allUpfront": "0.4897",
            "yrTerm1.allUpfront": "0.6124",
            "yrTerm1.partialUpfront": "0.6249",
            "yrTerm3.partialUpfront": "0.5209"
          },
          "ondemand": "0.953"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.392",
            "yrTerm3.allUpfront": "0.2246",
            "yrTerm1.allUpfront": "0.3329",
            "yrTerm1.partialUpfront": "0.3398",
            "yrTerm3.partialUpfront": "0.2389"
          },
          "ondemand": "0.529"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.87",
            "yrTerm3.allUpfront": "0.5736",
            "yrTerm1.allUpfront": "0.729",
            "yrTerm1.partialUpfront": "0.7438",
            "yrTerm3.partialUpfront": "0.6102"
          },
          "ondemand": "1.083"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.647",
            "yrTerm3.allUpfront": "1.1922",
            "yrTerm1.allUpfront": "1.3792",
            "yrTerm1.partialUpfront": "1.4073",
            "yrTerm3.partialUpfront": "1.2683"
          },
          "ondemand": "2.131"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.731",
            "yrTerm3.allUpfront": "0.4897",
            "yrTerm1.allUpfront": "0.6124",
            "yrTerm1.partialUpfront": "0.6249",
            "yrTerm3.partialUpfront": "0.5209"
          },
          "ondemand": "0.953"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.392",
            "yrTerm3.allUpfront": "0.2246",
            "yrTerm1.allUpfront": "0.3329",
            "yrTerm1.partialUpfront": "0.3398",
            "yrTerm3.partialUpfront": "0.2389"
          },
          "ondemand": "0.529"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.771",
            "yrTerm3.allUpfront": "0.5027",
            "yrTerm1.allUpfront": "0.6459",
            "yrTerm1.partialUpfront": "0.6591",
            "yrTerm3.partialUpfront": "0.5348"
          },
          "ondemand": "1.083"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.542",
            "yrTerm3.allUpfront": "1.1176",
            "yrTerm1.allUpfront": "1.2913",
            "yrTerm1.partialUpfront": "1.3176",
            "yrTerm3.partialUpfront": "1.1889"
          },
          "ondemand": "2.131"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.657",
            "yrTerm3.allUpfront": "0.437",
            "yrTerm1.allUpfront": "0.5507",
            "yrTerm1.partialUpfront": "0.5619",
            "yrTerm3.partialUpfront": "0.4649"
          },
          "ondemand": "0.752"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.292",
            "yrTerm3.allUpfront": "0.1572",
            "yrTerm1.allUpfront": "0.2477",
            "yrTerm1.partialUpfront": "0.2529",
            "yrTerm3.partialUpfront": "0.1673"
          },
          "ondemand": "0.42"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.831",
            "yrTerm3.allUpfront": "0.5338",
            "yrTerm1.allUpfront": "0.6961",
            "yrTerm1.partialUpfront": "0.7103",
            "yrTerm3.partialUpfront": "0.5679"
          },
          "ondemand": "1.167"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.602",
            "yrTerm3.allUpfront": "1.1499",
            "yrTerm1.allUpfront": "1.3414",
            "yrTerm1.partialUpfront": "1.3688",
            "yrTerm3.partialUpfront": "1.2233"
          },
          "ondemand": "2.215"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.717",
            "yrTerm3.allUpfront": "0.4686",
            "yrTerm1.allUpfront": "0.6005",
            "yrTerm1.partialUpfront": "0.6128",
            "yrTerm3.partialUpfront": "0.4985"
          },
          "ondemand": "0.836"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.352",
            "yrTerm3.allUpfront": "0.1887",
            "yrTerm1.allUpfront": "0.2973",
            "yrTerm1.partialUpfront": "0.3025",
            "yrTerm3.partialUpfront": "0.2007"
          },
          "ondemand": "0.504"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.771",
            "yrTerm3.allUpfront": "0.5027",
            "yrTerm1.allUpfront": "0.6459",
            "yrTerm1.partialUpfront": "0.6591",
            "yrTerm3.partialUpfront": "0.5348"
          },
          "ondemand": "1.083"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.542",
            "yrTerm3.allUpfront": "1.1176",
            "yrTerm1.allUpfront": "1.2913",
            "yrTerm1.partialUpfront": "1.3176",
            "yrTerm3.partialUpfront": "1.1889"
          },
          "ondemand": "2.131"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.741",
            "yrTerm3.allUpfront": "0.4963",
            "yrTerm1.allUpfront": "0.6203",
            "yrTerm1.partialUpfront": "0.6329",
            "yrTerm3.partialUpfront": "0.5279"
          },
          "ondemand": "0.752"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.36",
            "yrTerm3.allUpfront": "0.2061",
            "yrTerm1.allUpfront": "0.3075",
            "yrTerm1.partialUpfront": "0.3139",
            "yrTerm3.partialUpfront": "0.2193"
          },
          "ondemand": "0.478"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.872",
            "yrTerm3.allUpfront": "0.5672",
            "yrTerm1.allUpfront": "0.7306",
            "yrTerm1.partialUpfront": "0.7455",
            "yrTerm3.partialUpfront": "0.6034"
          },
          "ondemand": "1.713"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.387",
            "yrTerm3.allUpfront": "1.7243",
            "yrTerm1.allUpfront": "1.9993",
            "yrTerm1.partialUpfront": "2.0401",
            "yrTerm3.partialUpfront": "1.8343"
          },
          "ondemand": "3.421"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.758",
            "yrTerm3.allUpfront": "0.5021",
            "yrTerm1.allUpfront": "0.6353",
            "yrTerm1.partialUpfront": "0.6482",
            "yrTerm3.partialUpfront": "0.5342"
          },
          "ondemand": "0.848"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.384",
            "yrTerm3.allUpfront": "0.2221",
            "yrTerm1.allUpfront": "0.3321",
            "yrTerm1.partialUpfront": "0.3389",
            "yrTerm3.partialUpfront": "0.2363"
          },
          "ondemand": "0.516"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.771",
            "yrTerm3.allUpfront": "0.5027",
            "yrTerm1.allUpfront": "0.6459",
            "yrTerm1.partialUpfront": "0.6591",
            "yrTerm3.partialUpfront": "0.5348"
          },
          "ondemand": "1.083"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.542",
            "yrTerm3.allUpfront": "1.1176",
            "yrTerm1.allUpfront": "1.2913",
            "yrTerm1.partialUpfront": "1.3176",
            "yrTerm3.partialUpfront": "1.1889"
          },
          "ondemand": "2.131"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.741",
            "yrTerm3.allUpfront": "0.4963",
            "yrTerm1.allUpfront": "0.6203",
            "yrTerm1.partialUpfront": "0.6329",
            "yrTerm3.partialUpfront": "0.5279"
          },
          "ondemand": "0.752"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.36",
            "yrTerm3.allUpfront": "0.2061",
            "yrTerm1.allUpfront": "0.3075",
            "yrTerm1.partialUpfront": "0.3139",
            "yrTerm3.partialUpfront": "0.2193"
          },
          "ondemand": "0.478"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 15,
      "max_enis": 4
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "HVM",
      "PV"
    ],
    "ebs_optimized": true,
    "storage": {
      "ssd": true,
      "devices": 2,
      "size": 80.0
    },
    "instance_type": "c3.2xlarge",
    "ECU": 28.0,
    "memory": 15.0,
    "ebs_max_bandwidth": 1000.0
  },
  {
    "family": "Compute optimized",
    "enhanced_networking": true,
    "vCPU": 16,
    "generation": "current",
    "ebs_iops": 250.0,
    "network_performance": "High",
    "ebs_throughput": 16000.0,
    "pretty_name": "C3 High-CPU Quadruple Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.54",
            "yrTerm3.allUpfront": "1.0054",
            "yrTerm1.allUpfront": "1.2897",
            "yrTerm1.partialUpfront": "1.316",
            "yrTerm3.partialUpfront": "1.0695"
          },
          "ondemand": "2.166"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.084",
            "yrTerm3.allUpfront": "2.237",
            "yrTerm1.allUpfront": "2.5834",
            "yrTerm1.partialUpfront": "2.6361",
            "yrTerm3.partialUpfront": "2.3798"
          },
          "ondemand": "4.262"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.317",
            "yrTerm3.allUpfront": "0.8738",
            "yrTerm1.allUpfront": "1.1031",
            "yrTerm1.partialUpfront": "1.1256",
            "yrTerm3.partialUpfront": "0.9296"
          },
          "ondemand": "1.504"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.584",
            "yrTerm3.allUpfront": "0.3145",
            "yrTerm1.allUpfront": "0.4966",
            "yrTerm1.partialUpfront": "0.5067",
            "yrTerm3.partialUpfront": "0.3346"
          },
          "ondemand": "0.84"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.801",
            "yrTerm3.allUpfront": "1.2033",
            "yrTerm1.allUpfront": "1.5082",
            "yrTerm1.partialUpfront": "1.539",
            "yrTerm3.partialUpfront": "1.2801"
          },
          "ondemand": "2.166"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.493",
            "yrTerm3.allUpfront": "2.5554",
            "yrTerm1.allUpfront": "2.9253",
            "yrTerm1.partialUpfront": "2.9851",
            "yrTerm3.partialUpfront": "2.7185"
          },
          "ondemand": "4.564"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.489",
            "yrTerm3.allUpfront": "1.0164",
            "yrTerm1.allUpfront": "1.2474",
            "yrTerm1.partialUpfront": "1.2728",
            "yrTerm3.partialUpfront": "1.0813"
          },
          "ondemand": "1.849"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.776",
            "yrTerm3.allUpfront": "0.4634",
            "yrTerm1.allUpfront": "0.6868",
            "yrTerm1.partialUpfront": "0.7008",
            "yrTerm3.partialUpfront": "0.493"
          },
          "ondemand": "1.021"
        }
      },
      "sa-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "2.324",
            "yrTerm3.allUpfront": "1.2462",
            "yrTerm1.allUpfront": "1.9463",
            "yrTerm1.partialUpfront": "1.9861",
            "yrTerm3.partialUpfront": "1.3258"
          },
          "ondemand": "2.626"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "5.042",
            "yrTerm3.allUpfront": "2.7035",
            "yrTerm1.allUpfront": "4.2228",
            "yrTerm1.partialUpfront": "4.309",
            "yrTerm3.partialUpfront": "2.8761"
          },
          "ondemand": "5.174"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.839",
            "yrTerm3.allUpfront": "1.1146",
            "yrTerm1.allUpfront": "1.7403",
            "yrTerm1.partialUpfront": "1.7759",
            "yrTerm3.partialUpfront": "1.1858"
          },
          "ondemand": "1.964"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.984",
            "yrTerm3.allUpfront": "0.5538",
            "yrTerm1.allUpfront": "0.865",
            "yrTerm1.partialUpfront": "0.8827",
            "yrTerm3.partialUpfront": "0.5892"
          },
          "ondemand": "1.3"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.742",
            "yrTerm3.allUpfront": "1.1481",
            "yrTerm1.allUpfront": "1.4589",
            "yrTerm1.partialUpfront": "1.4887",
            "yrTerm3.partialUpfront": "1.2214"
          },
          "ondemand": "2.166"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.292",
            "yrTerm3.allUpfront": "2.3845",
            "yrTerm1.allUpfront": "2.7574",
            "yrTerm1.partialUpfront": "2.8137",
            "yrTerm3.partialUpfront": "2.5367"
          },
          "ondemand": "4.262"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.466",
            "yrTerm3.allUpfront": "0.9791",
            "yrTerm1.allUpfront": "1.2275",
            "yrTerm1.partialUpfront": "1.2526",
            "yrTerm3.partialUpfront": "1.0416"
          },
          "ondemand": "1.906"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.784",
            "yrTerm3.allUpfront": "0.4482",
            "yrTerm1.allUpfront": "0.6659",
            "yrTerm1.partialUpfront": "0.6796",
            "yrTerm3.partialUpfront": "0.4769"
          },
          "ondemand": "1.058"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.742",
            "yrTerm3.allUpfront": "1.1481",
            "yrTerm1.allUpfront": "1.4589",
            "yrTerm1.partialUpfront": "1.4887",
            "yrTerm3.partialUpfront": "1.2214"
          },
          "ondemand": "2.166"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.292",
            "yrTerm3.allUpfront": "2.3845",
            "yrTerm1.allUpfront": "2.7574",
            "yrTerm1.partialUpfront": "2.8137",
            "yrTerm3.partialUpfront": "2.5367"
          },
          "ondemand": "4.262"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.466",
            "yrTerm3.allUpfront": "0.9791",
            "yrTerm1.allUpfront": "1.2275",
            "yrTerm1.partialUpfront": "1.2526",
            "yrTerm3.partialUpfront": "1.0416"
          },
          "ondemand": "1.906"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.784",
            "yrTerm3.allUpfront": "0.4482",
            "yrTerm1.allUpfront": "0.6659",
            "yrTerm1.partialUpfront": "0.6796",
            "yrTerm3.partialUpfront": "0.4769"
          },
          "ondemand": "1.058"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.54",
            "yrTerm3.allUpfront": "1.0054",
            "yrTerm1.allUpfront": "1.2897",
            "yrTerm1.partialUpfront": "1.316",
            "yrTerm3.partialUpfront": "1.0695"
          },
          "ondemand": "2.166"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.084",
            "yrTerm3.allUpfront": "2.237",
            "yrTerm1.allUpfront": "2.5834",
            "yrTerm1.partialUpfront": "2.6361",
            "yrTerm3.partialUpfront": "2.3798"
          },
          "ondemand": "4.262"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.317",
            "yrTerm3.allUpfront": "0.8738",
            "yrTerm1.allUpfront": "1.1031",
            "yrTerm1.partialUpfront": "1.1256",
            "yrTerm3.partialUpfront": "0.9296"
          },
          "ondemand": "1.504"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.584",
            "yrTerm3.allUpfront": "0.3145",
            "yrTerm1.allUpfront": "0.4966",
            "yrTerm1.partialUpfront": "0.5067",
            "yrTerm3.partialUpfront": "0.3346"
          },
          "ondemand": "0.84"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.659",
            "yrTerm3.allUpfront": "1.0685",
            "yrTerm1.allUpfront": "1.3892",
            "yrTerm1.partialUpfront": "1.4176",
            "yrTerm3.partialUpfront": "1.1367"
          },
          "ondemand": "2.334"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.203",
            "yrTerm3.allUpfront": "2.3007",
            "yrTerm1.allUpfront": "2.6829",
            "yrTerm1.partialUpfront": "2.7377",
            "yrTerm3.partialUpfront": "2.4476"
          },
          "ondemand": "4.43"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.436",
            "yrTerm3.allUpfront": "0.937",
            "yrTerm1.allUpfront": "1.2027",
            "yrTerm1.partialUpfront": "1.2273",
            "yrTerm3.partialUpfront": "0.9968"
          },
          "ondemand": "1.672"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.704",
            "yrTerm3.allUpfront": "0.3774",
            "yrTerm1.allUpfront": "0.5959",
            "yrTerm1.partialUpfront": "0.6072",
            "yrTerm3.partialUpfront": "0.4015"
          },
          "ondemand": "1.008"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.54",
            "yrTerm3.allUpfront": "1.0054",
            "yrTerm1.allUpfront": "1.2897",
            "yrTerm1.partialUpfront": "1.316",
            "yrTerm3.partialUpfront": "1.0695"
          },
          "ondemand": "2.166"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.084",
            "yrTerm3.allUpfront": "2.237",
            "yrTerm1.allUpfront": "2.5834",
            "yrTerm1.partialUpfront": "2.6361",
            "yrTerm3.partialUpfront": "2.3798"
          },
          "ondemand": "4.262"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.482",
            "yrTerm3.allUpfront": "0.9913",
            "yrTerm1.allUpfront": "1.2412",
            "yrTerm1.partialUpfront": "1.2666",
            "yrTerm3.partialUpfront": "1.0546"
          },
          "ondemand": "1.504"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.72",
            "yrTerm3.allUpfront": "0.4123",
            "yrTerm1.allUpfront": "0.6151",
            "yrTerm1.partialUpfront": "0.6277",
            "yrTerm3.partialUpfront": "0.4386"
          },
          "ondemand": "0.956"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.74",
            "yrTerm3.allUpfront": "1.1366",
            "yrTerm1.allUpfront": "1.4576",
            "yrTerm1.partialUpfront": "1.4874",
            "yrTerm3.partialUpfront": "1.2091"
          },
          "ondemand": "2.383"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.991",
            "yrTerm3.allUpfront": "2.8859",
            "yrTerm1.allUpfront": "3.343",
            "yrTerm1.partialUpfront": "3.4113",
            "yrTerm3.partialUpfront": "3.0701"
          },
          "ondemand": "5.589"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.517",
            "yrTerm3.allUpfront": "1.0045",
            "yrTerm1.allUpfront": "1.2709",
            "yrTerm1.partialUpfront": "1.2968",
            "yrTerm3.partialUpfront": "1.0686"
          },
          "ondemand": "1.696"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.768",
            "yrTerm3.allUpfront": "0.4452",
            "yrTerm1.allUpfront": "0.6643",
            "yrTerm1.partialUpfront": "0.6778",
            "yrTerm3.partialUpfront": "0.4737"
          },
          "ondemand": "1.032"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.54",
            "yrTerm3.allUpfront": "1.0054",
            "yrTerm1.allUpfront": "1.2897",
            "yrTerm1.partialUpfront": "1.316",
            "yrTerm3.partialUpfront": "1.0695"
          },
          "ondemand": "2.166"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.084",
            "yrTerm3.allUpfront": "2.237",
            "yrTerm1.allUpfront": "2.5834",
            "yrTerm1.partialUpfront": "2.6361",
            "yrTerm3.partialUpfront": "2.3798"
          },
          "ondemand": "4.262"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.482",
            "yrTerm3.allUpfront": "0.9913",
            "yrTerm1.allUpfront": "1.2412",
            "yrTerm1.partialUpfront": "1.2666",
            "yrTerm3.partialUpfront": "1.0546"
          },
          "ondemand": "1.504"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.72",
            "yrTerm3.allUpfront": "0.4123",
            "yrTerm1.allUpfront": "0.6151",
            "yrTerm1.partialUpfront": "0.6277",
            "yrTerm3.partialUpfront": "0.4386"
          },
          "ondemand": "0.956"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 30,
      "max_enis": 8
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "HVM",
      "PV"
    ],
    "ebs_optimized": true,
    "storage": {
      "ssd": true,
      "devices": 2,
      "size": 160.0
    },
    "instance_type": "c3.4xlarge",
    "ECU": 55.0,
    "memory": 30.0,
    "ebs_max_bandwidth": 2000.0
  },
  {
    "family": "Compute optimized",
    "enhanced_networking": true,
    "vCPU": 32,
    "generation": "current",
    "ebs_iops": 0,
    "network_performance": "10 Gigabit",
    "ebs_throughput": 0,
    "pretty_name": "C3 High-CPU Eight Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.081",
            "yrTerm3.allUpfront": "2.0098",
            "yrTerm1.allUpfront": "2.5805",
            "yrTerm1.partialUpfront": "2.6332",
            "yrTerm3.partialUpfront": "2.1381"
          },
          "ondemand": "4.332"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "6.17",
            "yrTerm3.allUpfront": "4.4741",
            "yrTerm1.allUpfront": "5.1677",
            "yrTerm1.partialUpfront": "5.2732",
            "yrTerm3.partialUpfront": "4.7597"
          },
          "ondemand": "8.523"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.631",
            "yrTerm3.allUpfront": "1.7455",
            "yrTerm1.allUpfront": "2.204",
            "yrTerm1.partialUpfront": "2.249",
            "yrTerm3.partialUpfront": "1.857"
          },
          "ondemand": "3.008"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.168",
            "yrTerm3.allUpfront": "0.6281",
            "yrTerm1.allUpfront": "0.9921",
            "yrTerm1.partialUpfront": "1.0124",
            "yrTerm3.partialUpfront": "0.6682"
          },
          "ondemand": "1.68"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.599",
            "yrTerm3.allUpfront": "2.4047",
            "yrTerm1.allUpfront": "3.0143",
            "yrTerm1.partialUpfront": "3.0758",
            "yrTerm3.partialUpfront": "2.5582"
          },
          "ondemand": "4.332"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "6.986",
            "yrTerm3.allUpfront": "5.1117",
            "yrTerm1.allUpfront": "5.8518",
            "yrTerm1.partialUpfront": "5.9713",
            "yrTerm3.partialUpfront": "5.438"
          },
          "ondemand": "9.128"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.978",
            "yrTerm3.allUpfront": "2.034",
            "yrTerm1.allUpfront": "2.4942",
            "yrTerm1.partialUpfront": "2.545",
            "yrTerm3.partialUpfront": "2.1638"
          },
          "ondemand": "3.699"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.552",
            "yrTerm3.allUpfront": "0.9259",
            "yrTerm1.allUpfront": "1.3735",
            "yrTerm1.partialUpfront": "1.4016",
            "yrTerm3.partialUpfront": "0.985"
          },
          "ondemand": "2.043"
        }
      },
      "sa-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "4.648",
            "yrTerm3.allUpfront": "2.4924",
            "yrTerm1.allUpfront": "3.8928",
            "yrTerm1.partialUpfront": "3.9722",
            "yrTerm3.partialUpfront": "2.6515"
          },
          "ondemand": "5.252"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "10.085",
            "yrTerm3.allUpfront": "5.4088",
            "yrTerm1.allUpfront": "8.4476",
            "yrTerm1.partialUpfront": "8.62",
            "yrTerm3.partialUpfront": "5.754"
          },
          "ondemand": "10.347"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "3.676",
            "yrTerm3.allUpfront": "2.2274",
            "yrTerm1.allUpfront": "3.4796",
            "yrTerm1.partialUpfront": "3.5506",
            "yrTerm3.partialUpfront": "2.3696"
          },
          "ondemand": "3.928"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.968",
            "yrTerm3.allUpfront": "1.1077",
            "yrTerm1.allUpfront": "1.73",
            "yrTerm1.partialUpfront": "1.7654",
            "yrTerm3.partialUpfront": "1.1784"
          },
          "ondemand": "2.6"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.485",
            "yrTerm3.allUpfront": "2.2963",
            "yrTerm1.allUpfront": "2.9187",
            "yrTerm1.partialUpfront": "2.9783",
            "yrTerm3.partialUpfront": "2.4429"
          },
          "ondemand": "4.332"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "6.584",
            "yrTerm3.allUpfront": "4.7689",
            "yrTerm1.allUpfront": "5.5147",
            "yrTerm1.partialUpfront": "5.6273",
            "yrTerm3.partialUpfront": "5.0733"
          },
          "ondemand": "8.523"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.926",
            "yrTerm3.allUpfront": "1.9571",
            "yrTerm1.allUpfront": "2.451",
            "yrTerm1.partialUpfront": "2.501",
            "yrTerm3.partialUpfront": "2.082"
          },
          "ondemand": "3.813"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.568",
            "yrTerm3.allUpfront": "0.8965",
            "yrTerm1.allUpfront": "1.3309",
            "yrTerm1.partialUpfront": "1.3581",
            "yrTerm3.partialUpfront": "0.9537"
          },
          "ondemand": "2.117"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.485",
            "yrTerm3.allUpfront": "2.2963",
            "yrTerm1.allUpfront": "2.9187",
            "yrTerm1.partialUpfront": "2.9783",
            "yrTerm3.partialUpfront": "2.4429"
          },
          "ondemand": "4.332"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "6.584",
            "yrTerm3.allUpfront": "4.7689",
            "yrTerm1.allUpfront": "5.5147",
            "yrTerm1.partialUpfront": "5.6273",
            "yrTerm3.partialUpfront": "5.0733"
          },
          "ondemand": "8.523"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.926",
            "yrTerm3.allUpfront": "1.9571",
            "yrTerm1.allUpfront": "2.451",
            "yrTerm1.partialUpfront": "2.501",
            "yrTerm3.partialUpfront": "2.082"
          },
          "ondemand": "3.813"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.568",
            "yrTerm3.allUpfront": "0.8965",
            "yrTerm1.allUpfront": "1.3309",
            "yrTerm1.partialUpfront": "1.3581",
            "yrTerm3.partialUpfront": "0.9537"
          },
          "ondemand": "2.117"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.081",
            "yrTerm3.allUpfront": "2.0098",
            "yrTerm1.allUpfront": "2.5805",
            "yrTerm1.partialUpfront": "2.6332",
            "yrTerm3.partialUpfront": "2.1381"
          },
          "ondemand": "4.332"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "6.17",
            "yrTerm3.allUpfront": "4.4741",
            "yrTerm1.allUpfront": "5.1677",
            "yrTerm1.partialUpfront": "5.2732",
            "yrTerm3.partialUpfront": "4.7597"
          },
          "ondemand": "8.523"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.631",
            "yrTerm3.allUpfront": "1.7455",
            "yrTerm1.allUpfront": "2.204",
            "yrTerm1.partialUpfront": "2.249",
            "yrTerm3.partialUpfront": "1.857"
          },
          "ondemand": "3.008"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.168",
            "yrTerm3.allUpfront": "0.6281",
            "yrTerm1.allUpfront": "0.9921",
            "yrTerm1.partialUpfront": "1.0124",
            "yrTerm3.partialUpfront": "0.6682"
          },
          "ondemand": "1.68"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.318",
            "yrTerm3.allUpfront": "2.1361",
            "yrTerm1.allUpfront": "2.7793",
            "yrTerm1.partialUpfront": "2.8361",
            "yrTerm3.partialUpfront": "2.2725"
          },
          "ondemand": "4.668"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "6.407",
            "yrTerm3.allUpfront": "4.6005",
            "yrTerm1.allUpfront": "5.3667",
            "yrTerm1.partialUpfront": "5.4762",
            "yrTerm3.partialUpfront": "4.8941"
          },
          "ondemand": "8.859"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.869",
            "yrTerm3.allUpfront": "1.8719",
            "yrTerm1.allUpfront": "2.4032",
            "yrTerm1.partialUpfront": "2.4524",
            "yrTerm3.partialUpfront": "1.9914"
          },
          "ondemand": "3.344"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.408",
            "yrTerm3.allUpfront": "0.7537",
            "yrTerm1.allUpfront": "1.1905",
            "yrTerm1.partialUpfront": "1.2145",
            "yrTerm3.partialUpfront": "0.801"
          },
          "ondemand": "2.016"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.081",
            "yrTerm3.allUpfront": "2.0098",
            "yrTerm1.allUpfront": "2.5805",
            "yrTerm1.partialUpfront": "2.6332",
            "yrTerm3.partialUpfront": "2.1381"
          },
          "ondemand": "4.332"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "6.17",
            "yrTerm3.allUpfront": "4.4741",
            "yrTerm1.allUpfront": "5.1677",
            "yrTerm1.partialUpfront": "5.2732",
            "yrTerm3.partialUpfront": "4.7597"
          },
          "ondemand": "8.523"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.961",
            "yrTerm3.allUpfront": "1.9815",
            "yrTerm1.allUpfront": "2.4804",
            "yrTerm1.partialUpfront": "2.531",
            "yrTerm3.partialUpfront": "2.108"
          },
          "ondemand": "3.008"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.44",
            "yrTerm3.allUpfront": "0.8245",
            "yrTerm1.allUpfront": "1.2313",
            "yrTerm1.partialUpfront": "1.2564",
            "yrTerm3.partialUpfront": "0.8772"
          },
          "ondemand": "1.912"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.484",
            "yrTerm3.allUpfront": "2.2722",
            "yrTerm1.allUpfront": "2.9183",
            "yrTerm1.partialUpfront": "2.9778",
            "yrTerm3.partialUpfront": "2.4172"
          },
          "ondemand": "4.716"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "7.496",
            "yrTerm3.allUpfront": "5.422",
            "yrTerm1.allUpfront": "6.2787",
            "yrTerm1.partialUpfront": "6.4068",
            "yrTerm3.partialUpfront": "5.7681"
          },
          "ondemand": "10.401"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "3.034",
            "yrTerm3.allUpfront": "2.0077",
            "yrTerm1.allUpfront": "2.5414",
            "yrTerm1.partialUpfront": "2.5934",
            "yrTerm3.partialUpfront": "2.1358"
          },
          "ondemand": "3.392"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.536",
            "yrTerm3.allUpfront": "0.8905",
            "yrTerm1.allUpfront": "1.3297",
            "yrTerm1.partialUpfront": "1.3569",
            "yrTerm3.partialUpfront": "0.9474"
          },
          "ondemand": "2.064"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.081",
            "yrTerm3.allUpfront": "2.0098",
            "yrTerm1.allUpfront": "2.5805",
            "yrTerm1.partialUpfront": "2.6332",
            "yrTerm3.partialUpfront": "2.1381"
          },
          "ondemand": "4.332"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "6.17",
            "yrTerm3.allUpfront": "4.4741",
            "yrTerm1.allUpfront": "5.1677",
            "yrTerm1.partialUpfront": "5.2732",
            "yrTerm3.partialUpfront": "4.7597"
          },
          "ondemand": "8.523"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.961",
            "yrTerm3.allUpfront": "1.9815",
            "yrTerm1.allUpfront": "2.4804",
            "yrTerm1.partialUpfront": "2.531",
            "yrTerm3.partialUpfront": "2.108"
          },
          "ondemand": "3.008"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.44",
            "yrTerm3.allUpfront": "0.8245",
            "yrTerm1.allUpfront": "1.2313",
            "yrTerm1.partialUpfront": "1.2564",
            "yrTerm3.partialUpfront": "0.8772"
          },
          "ondemand": "1.912"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 30,
      "max_enis": 8
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "HVM",
      "PV"
    ],
    "ebs_optimized": false,
    "storage": {
      "ssd": true,
      "devices": 2,
      "size": 320.0
    },
    "instance_type": "c3.8xlarge",
    "ECU": 108.0,
    "memory": 60.0,
    "ebs_max_bandwidth": 0
  },
  {
    "family": "GPU instances",
    "enhanced_networking": false,
    "vCPU": 8,
    "generation": "current",
    "ebs_iops": 125.0,
    "network_performance": "High",
    "ebs_throughput": 8000.0,
    "pretty_name": "G2 Double Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.775",
            "yrTerm3.allUpfront": "0.4609",
            "yrTerm1.allUpfront": "0.6493",
            "yrTerm1.partialUpfront": "0.6626",
            "yrTerm3.partialUpfront": "0.4903"
          },
          "ondemand": "0.957"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.754",
            "yrTerm3.allUpfront": "2.7072",
            "yrTerm1.allUpfront": "3.1441",
            "yrTerm1.partialUpfront": "3.2082",
            "yrTerm3.partialUpfront": "2.88"
          },
          "ondemand": "3.816"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.611",
            "yrTerm3.allUpfront": "0.3923",
            "yrTerm1.allUpfront": "0.5121",
            "yrTerm1.partialUpfront": "0.5225",
            "yrTerm3.partialUpfront": "0.4173"
          },
          "ondemand": "0.767"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.474",
            "yrTerm3.allUpfront": "0.282",
            "yrTerm1.allUpfront": "0.397",
            "yrTerm1.partialUpfront": "0.4052",
            "yrTerm3.partialUpfront": "0.3"
          },
          "ondemand": "0.65"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.986",
            "yrTerm3.allUpfront": "0.5871",
            "yrTerm1.allUpfront": "0.8261",
            "yrTerm1.partialUpfront": "0.8431",
            "yrTerm3.partialUpfront": "0.6246"
          },
          "ondemand": "1.19"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.824",
            "yrTerm3.allUpfront": "2.726",
            "yrTerm1.allUpfront": "3.2027",
            "yrTerm1.partialUpfront": "3.2681",
            "yrTerm3.partialUpfront": "2.9"
          },
          "ondemand": "3.913"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.831",
            "yrTerm3.allUpfront": "0.5213",
            "yrTerm1.allUpfront": "0.6961",
            "yrTerm1.partialUpfront": "0.7103",
            "yrTerm3.partialUpfront": "0.5546"
          },
          "ondemand": "1.01"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.678",
            "yrTerm3.allUpfront": "0.4166",
            "yrTerm1.allUpfront": "0.5868",
            "yrTerm1.partialUpfront": "0.5988",
            "yrTerm3.partialUpfront": "0.4433"
          },
          "ondemand": "0.898"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.162",
            "yrTerm3.allUpfront": "0.6435",
            "yrTerm1.allUpfront": "0.9735",
            "yrTerm1.partialUpfront": "0.9934",
            "yrTerm3.partialUpfront": "0.6846"
          },
          "ondemand": "1.307"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "4.064",
            "yrTerm3.allUpfront": "2.8892",
            "yrTerm1.allUpfront": "3.933",
            "yrTerm1.partialUpfront": "4.0133",
            "yrTerm3.partialUpfront": "3.0735"
          },
          "ondemand": "4.166"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.036",
            "yrTerm3.allUpfront": "0.5741",
            "yrTerm1.allUpfront": "0.8679",
            "yrTerm1.partialUpfront": "0.8857",
            "yrTerm3.partialUpfront": "0.6107"
          },
          "ondemand": "1.16"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.772",
            "yrTerm3.allUpfront": "0.4642",
            "yrTerm1.allUpfront": "0.7021",
            "yrTerm1.partialUpfront": "0.7164",
            "yrTerm3.partialUpfront": "0.4938"
          },
          "ondemand": "1"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.002",
            "yrTerm3.allUpfront": "0.5959",
            "yrTerm1.allUpfront": "0.8393",
            "yrTerm1.partialUpfront": "0.8564",
            "yrTerm3.partialUpfront": "0.6339"
          },
          "ondemand": "1.205"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.996",
            "yrTerm3.allUpfront": "2.8419",
            "yrTerm1.allUpfront": "3.8822",
            "yrTerm1.partialUpfront": "3.9614",
            "yrTerm3.partialUpfront": "3.0232"
          },
          "ondemand": "4.064"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.885",
            "yrTerm3.allUpfront": "0.5262",
            "yrTerm1.allUpfront": "0.7413",
            "yrTerm1.partialUpfront": "0.7564",
            "yrTerm3.partialUpfront": "0.5598"
          },
          "ondemand": "1.058"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.678",
            "yrTerm3.allUpfront": "0.4166",
            "yrTerm1.allUpfront": "0.5868",
            "yrTerm1.partialUpfront": "0.5988",
            "yrTerm3.partialUpfront": "0.4433"
          },
          "ondemand": "0.898"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.775",
            "yrTerm3.allUpfront": "0.4609",
            "yrTerm1.allUpfront": "0.6493",
            "yrTerm1.partialUpfront": "0.6626",
            "yrTerm3.partialUpfront": "0.4903"
          },
          "ondemand": "0.957"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.754",
            "yrTerm3.allUpfront": "2.7072",
            "yrTerm1.allUpfront": "3.1441",
            "yrTerm1.partialUpfront": "3.2082",
            "yrTerm3.partialUpfront": "2.88"
          },
          "ondemand": "3.816"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.611",
            "yrTerm3.allUpfront": "0.3923",
            "yrTerm1.allUpfront": "0.5121",
            "yrTerm1.partialUpfront": "0.5225",
            "yrTerm3.partialUpfront": "0.4173"
          },
          "ondemand": "0.767"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.474",
            "yrTerm3.allUpfront": "0.282",
            "yrTerm1.allUpfront": "0.397",
            "yrTerm1.partialUpfront": "0.4052",
            "yrTerm3.partialUpfront": "0.3"
          },
          "ondemand": "0.65"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.816",
            "yrTerm3.allUpfront": "0.4847",
            "yrTerm1.allUpfront": "0.6832",
            "yrTerm1.partialUpfront": "0.6972",
            "yrTerm3.partialUpfront": "0.5156"
          },
          "ondemand": "1.009"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.793",
            "yrTerm3.allUpfront": "2.73",
            "yrTerm1.allUpfront": "3.1768",
            "yrTerm1.partialUpfront": "3.2417",
            "yrTerm3.partialUpfront": "2.9042"
          },
          "ondemand": "3.868"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.651",
            "yrTerm3.allUpfront": "0.4156",
            "yrTerm1.allUpfront": "0.5452",
            "yrTerm1.partialUpfront": "0.5563",
            "yrTerm3.partialUpfront": "0.4421"
          },
          "ondemand": "0.819"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.514",
            "yrTerm3.allUpfront": "0.305",
            "yrTerm1.allUpfront": "0.4306",
            "yrTerm1.partialUpfront": "0.4394",
            "yrTerm3.partialUpfront": "0.3245"
          },
          "ondemand": "0.702"
        }
      },
      "eu-central-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.649",
            "yrTerm3.allUpfront": "0.4023",
            "yrTerm1.allUpfront": "0.5433",
            "yrTerm1.partialUpfront": "0.5542",
            "yrTerm3.partialUpfront": "0.428"
          },
          "ondemand": "0.906"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.565",
            "yrTerm3.allUpfront": "0.3355",
            "yrTerm1.allUpfront": "0.4736",
            "yrTerm1.partialUpfront": "0.4837",
            "yrTerm3.partialUpfront": "0.3565"
          },
          "ondemand": "0.772"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.775",
            "yrTerm3.allUpfront": "0.4609",
            "yrTerm1.allUpfront": "0.6493",
            "yrTerm1.partialUpfront": "0.6626",
            "yrTerm3.partialUpfront": "0.4903"
          },
          "ondemand": "0.957"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.754",
            "yrTerm3.allUpfront": "2.7072",
            "yrTerm1.allUpfront": "3.1441",
            "yrTerm1.partialUpfront": "3.2082",
            "yrTerm3.partialUpfront": "2.88"
          },
          "ondemand": "3.816"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.611",
            "yrTerm3.allUpfront": "0.3894",
            "yrTerm1.allUpfront": "0.5121",
            "yrTerm1.partialUpfront": "0.5225",
            "yrTerm3.partialUpfront": "0.4143"
          },
          "ondemand": "0.767"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.514",
            "yrTerm3.allUpfront": "0.305",
            "yrTerm1.allUpfront": "0.4306",
            "yrTerm1.partialUpfront": "0.4394",
            "yrTerm3.partialUpfront": "0.3245"
          },
          "ondemand": "0.702"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 15,
      "max_enis": 4
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": true,
    "storage": {
      "ssd": true,
      "devices": 1,
      "size": 60.0
    },
    "instance_type": "g2.2xlarge",
    "ECU": 26.0,
    "memory": 15.0,
    "ebs_max_bandwidth": 1000.0
  },
  {
    "family": "GPU instances",
    "enhanced_networking": false,
    "vCPU": 32,
    "generation": "current",
    "ebs_iops": 0,
    "network_performance": "10 Gigabit",
    "ebs_throughput": 0,
    "pretty_name": "G2 Eight Extra Large",
    "pricing": {
      "us-east-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.9791",
            "yrTerm3.allUpfront": "1.194",
            "yrTerm1.allUpfront": "1.6576",
            "yrTerm1.partialUpfront": "1.6918",
            "yrTerm3.partialUpfront": "1.271"
          },
          "ondemand": "2.878"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.896",
            "yrTerm3.allUpfront": "1.1279",
            "yrTerm1.allUpfront": "1.5881",
            "yrTerm1.partialUpfront": "1.621",
            "yrTerm3.partialUpfront": "1.2"
          },
          "ondemand": "2.6"
        }
      },
      "ap-northeast-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.7951",
            "yrTerm3.allUpfront": "1.7324",
            "yrTerm1.allUpfront": "2.4168",
            "yrTerm1.partialUpfront": "2.4662",
            "yrTerm3.partialUpfront": "1.8442"
          },
          "ondemand": "3.87"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.712",
            "yrTerm3.allUpfront": "1.6665",
            "yrTerm1.allUpfront": "2.347",
            "yrTerm1.partialUpfront": "2.3953",
            "yrTerm3.partialUpfront": "1.773"
          },
          "ondemand": "3.592"
        }
      },
      "ap-southeast-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "3.1711",
            "yrTerm3.allUpfront": "1.9228",
            "yrTerm1.allUpfront": "2.878",
            "yrTerm1.partialUpfront": "2.9366",
            "yrTerm3.partialUpfront": "2.0462"
          },
          "ondemand": "4.278"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "3.088",
            "yrTerm3.allUpfront": "1.8566",
            "yrTerm1.allUpfront": "2.8082",
            "yrTerm1.partialUpfront": "2.8656",
            "yrTerm3.partialUpfront": "1.9753"
          },
          "ondemand": "4"
        }
      },
      "ap-southeast-2": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.7951",
            "yrTerm3.allUpfront": "1.7324",
            "yrTerm1.allUpfront": "2.4168",
            "yrTerm1.partialUpfront": "2.4662",
            "yrTerm3.partialUpfront": "1.8442"
          },
          "ondemand": "3.87"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.712",
            "yrTerm3.allUpfront": "1.6665",
            "yrTerm1.allUpfront": "2.347",
            "yrTerm1.partialUpfront": "2.3953",
            "yrTerm3.partialUpfront": "1.773"
          },
          "ondemand": "3.592"
        }
      },
      "us-west-2": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.9791",
            "yrTerm3.allUpfront": "1.194",
            "yrTerm1.allUpfront": "1.6576",
            "yrTerm1.partialUpfront": "1.6918",
            "yrTerm3.partialUpfront": "1.271"
          },
          "ondemand": "2.878"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.896",
            "yrTerm3.allUpfront": "1.1279",
            "yrTerm1.allUpfront": "1.5881",
            "yrTerm1.partialUpfront": "1.621",
            "yrTerm3.partialUpfront": "1.2"
          },
          "ondemand": "2.6"
        }
      },
      "us-west-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.1391",
            "yrTerm3.allUpfront": "1.286",
            "yrTerm1.allUpfront": "1.792",
            "yrTerm1.partialUpfront": "1.8286",
            "yrTerm3.partialUpfront": "1.369"
          },
          "ondemand": "3.086"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.056",
            "yrTerm3.allUpfront": "1.2201",
            "yrTerm1.allUpfront": "1.7224",
            "yrTerm1.partialUpfront": "1.7576",
            "yrTerm3.partialUpfront": "1.2981"
          },
          "ondemand": "2.808"
        }
      },
      "eu-central-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.347",
            "yrTerm3.allUpfront": "1.4072",
            "yrTerm1.allUpfront": "1.9655",
            "yrTerm1.partialUpfront": "2.0059",
            "yrTerm3.partialUpfront": "1.4975"
          },
          "ondemand": "3.366"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.26",
            "yrTerm3.allUpfront": "1.3422",
            "yrTerm1.allUpfront": "1.8945",
            "yrTerm1.partialUpfront": "1.9347",
            "yrTerm3.partialUpfront": "1.426"
          },
          "ondemand": "3.088"
        }
      },
      "eu-west-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.1391",
            "yrTerm3.allUpfront": "1.286",
            "yrTerm1.allUpfront": "1.792",
            "yrTerm1.partialUpfront": "1.8286",
            "yrTerm3.partialUpfront": "1.369"
          },
          "ondemand": "3.086"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.056",
            "yrTerm3.allUpfront": "1.2201",
            "yrTerm1.allUpfront": "1.7224",
            "yrTerm1.partialUpfront": "1.7576",
            "yrTerm3.partialUpfront": "1.2981"
          },
          "ondemand": "2.808"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 30,
      "max_enis": 8
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": false,
    "storage": {
      "ssd": true,
      "devices": 2,
      "size": 120.0
    },
    "instance_type": "g2.8xlarge",
    "ECU": 104.0,
    "memory": 60.0,
    "ebs_max_bandwidth": 0
  },
  {
    "family": "Memory optimized",
    "enhanced_networking": false,
    "vCPU": 128,
    "generation": "current",
    "ebs_iops": 0,
    "network_performance": "10 Gigabit",
    "ebs_throughput": 0,
    "pretty_name": "X1 32xlarge",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "17.700",
            "yrTerm3.allUpfront": "12.274",
            "yrTerm1.allUpfront": "16.213",
            "yrTerm1.partialUpfront": "16.370",
            "yrTerm3.partialUpfront": "12.512"
          },
          "ondemand": "21.880"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "57.502",
            "yrTerm3.allUpfront": "52.076",
            "yrTerm1.allUpfront": "56.015",
            "yrTerm1.partialUpfront": "56.172",
            "yrTerm3.partialUpfront": "52.314"
          },
          "ondemand": "61.682"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "15.046",
            "yrTerm3.allUpfront": "9.620",
            "yrTerm1.allUpfront": "13.559",
            "yrTerm1.partialUpfront": "13.716",
            "yrTerm3.partialUpfront": "9.858"
          },
          "ondemand": "19.226"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "9.158",
            "yrTerm3.allUpfront": "3.732",
            "yrTerm1.allUpfront": "7.671",
            "yrTerm1.partialUpfront": "7.828",
            "yrTerm3.partialUpfront": "3.970"
          },
          "ondemand": "13.338"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "21.822",
            "yrTerm3.allUpfront": "13.953",
            "yrTerm1.allUpfront": "19.665",
            "yrTerm1.partialUpfront": "19.892",
            "yrTerm3.partialUpfront": "14.299"
          },
          "ondemand": "27.883"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "61.624",
            "yrTerm3.allUpfront": "53.755",
            "yrTerm1.allUpfront": "59.467",
            "yrTerm1.partialUpfront": "59.694",
            "yrTerm3.partialUpfront": "54.101"
          },
          "ondemand": "67.685"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "19.168",
            "yrTerm3.allUpfront": "11.299",
            "yrTerm1.allUpfront": "17.011",
            "yrTerm1.partialUpfront": "17.238",
            "yrTerm3.partialUpfront": "11.645"
          },
          "ondemand": "25.229"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "13.280",
            "yrTerm3.allUpfront": "5.411",
            "yrTerm1.allUpfront": "11.123",
            "yrTerm1.partialUpfront": "11.350",
            "yrTerm3.partialUpfront": "5.757"
          },
          "ondemand": "19.341"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "21.822",
            "yrTerm3.allUpfront": "13.953",
            "yrTerm1.allUpfront": "19.665",
            "yrTerm1.partialUpfront": "19.892",
            "yrTerm3.partialUpfront": "14.299"
          },
          "ondemand": "27.883"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "61.624",
            "yrTerm3.allUpfront": "53.755",
            "yrTerm1.allUpfront": "59.467",
            "yrTerm1.partialUpfront": "59.694",
            "yrTerm3.partialUpfront": "54.101"
          },
          "ondemand": "67.685"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "19.168",
            "yrTerm3.allUpfront": "11.299",
            "yrTerm1.allUpfront": "17.011",
            "yrTerm1.partialUpfront": "17.238",
            "yrTerm3.partialUpfront": "11.645"
          },
          "ondemand": "25.229"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "13.280",
            "yrTerm3.allUpfront": "5.411",
            "yrTerm1.allUpfront": "11.123",
            "yrTerm1.partialUpfront": "11.350",
            "yrTerm3.partialUpfront": "5.757"
          },
          "ondemand": "19.341"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "21.822",
            "yrTerm3.allUpfront": "13.953",
            "yrTerm1.allUpfront": "19.665",
            "yrTerm1.partialUpfront": "19.892",
            "yrTerm3.partialUpfront": "14.299"
          },
          "ondemand": "27.883"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "61.624",
            "yrTerm3.allUpfront": "53.755",
            "yrTerm1.allUpfront": "59.467",
            "yrTerm1.partialUpfront": "59.694",
            "yrTerm3.partialUpfront": "54.101"
          },
          "ondemand": "67.685"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "19.168",
            "yrTerm3.allUpfront": "11.299",
            "yrTerm1.allUpfront": "17.011",
            "yrTerm1.partialUpfront": "17.238",
            "yrTerm3.partialUpfront": "11.645"
          },
          "ondemand": "25.229"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "13.280",
            "yrTerm3.allUpfront": "5.411",
            "yrTerm1.allUpfront": "11.123",
            "yrTerm1.partialUpfront": "11.350",
            "yrTerm3.partialUpfront": "5.757"
          },
          "ondemand": "19.341"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "17.700",
            "yrTerm3.allUpfront": "12.274",
            "yrTerm1.allUpfront": "16.213",
            "yrTerm1.partialUpfront": "16.370",
            "yrTerm3.partialUpfront": "12.512"
          },
          "ondemand": "21.880"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "57.502",
            "yrTerm3.allUpfront": "52.076",
            "yrTerm1.allUpfront": "56.015",
            "yrTerm1.partialUpfront": "56.172",
            "yrTerm3.partialUpfront": "52.314"
          },
          "ondemand": "61.682"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "15.046",
            "yrTerm3.allUpfront": "9.620",
            "yrTerm1.allUpfront": "13.559",
            "yrTerm1.partialUpfront": "13.716",
            "yrTerm3.partialUpfront": "9.858"
          },
          "ondemand": "19.226"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "9.158",
            "yrTerm3.allUpfront": "3.732",
            "yrTerm1.allUpfront": "7.671",
            "yrTerm1.partialUpfront": "7.828",
            "yrTerm3.partialUpfront": "3.970"
          },
          "ondemand": "13.338"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "21.364",
            "yrTerm3.allUpfront": "13.767",
            "yrTerm1.allUpfront": "19.282",
            "yrTerm1.partialUpfront": "19.501",
            "yrTerm3.partialUpfront": "14.100"
          },
          "ondemand": "27.216"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "61.166",
            "yrTerm3.allUpfront": "53.569",
            "yrTerm1.allUpfront": "59.084",
            "yrTerm1.partialUpfront": "59.303",
            "yrTerm3.partialUpfront": "53.902"
          },
          "ondemand": "67.018"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "18.710",
            "yrTerm3.allUpfront": "11.113",
            "yrTerm1.allUpfront": "16.628",
            "yrTerm1.partialUpfront": "16.847",
            "yrTerm3.partialUpfront": "11.446"
          },
          "ondemand": "24.562"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "12.822",
            "yrTerm3.allUpfront": "5.225",
            "yrTerm1.allUpfront": "10.740",
            "yrTerm1.partialUpfront": "10.959",
            "yrTerm3.partialUpfront": "5.558"
          },
          "ondemand": "18.674"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "19.532",
            "yrTerm3.allUpfront": "13.020",
            "yrTerm1.allUpfront": "17.747",
            "yrTerm1.partialUpfront": "17.935",
            "yrTerm3.partialUpfront": "13.306"
          },
          "ondemand": "24.548"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "59.334",
            "yrTerm3.allUpfront": "52.822",
            "yrTerm1.allUpfront": "57.549",
            "yrTerm1.partialUpfront": "57.737",
            "yrTerm3.partialUpfront": "53.108"
          },
          "ondemand": "64.350"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "16.878",
            "yrTerm3.allUpfront": "10.366",
            "yrTerm1.allUpfront": "15.093",
            "yrTerm1.partialUpfront": "15.281",
            "yrTerm3.partialUpfront": "10.652"
          },
          "ondemand": "21.894"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "10.990",
            "yrTerm3.allUpfront": "4.478",
            "yrTerm1.allUpfront": "9.205",
            "yrTerm1.partialUpfront": "9.393",
            "yrTerm3.partialUpfront": "4.764"
          },
          "ondemand": "16.006"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 30,
      "max_enis": 8
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": true,
    "storage": {
      "ssd": true,
      "devices": 2,
      "size": 1920.0
    },
    "instance_type": "x1.32xlarge",
    "ECU": 349.0,
    "memory": 1952.0,
    "ebs_max_bandwidth": 0
  },
  {
    "family": "Memory optimized",
    "enhanced_networking": true,
    "vCPU": 2,
    "generation": "current",
    "ebs_iops": 0,
    "network_performance": "Moderate",
    "ebs_throughput": 0,
    "pretty_name": "R3 High-Memory Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.272",
            "yrTerm3.allUpfront": "0.1899",
            "yrTerm1.allUpfront": "0.2276",
            "yrTerm1.partialUpfront": "0.2319",
            "yrTerm3.partialUpfront": "0.2023"
          },
          "ondemand": "0.386"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.833",
            "yrTerm3.allUpfront": "0.5985",
            "yrTerm1.allUpfront": "0.6974",
            "yrTerm1.partialUpfront": "0.7113",
            "yrTerm3.partialUpfront": "0.637"
          },
          "ondemand": "0.946"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.238",
            "yrTerm3.allUpfront": "0.1457",
            "yrTerm1.allUpfront": "0.1992",
            "yrTerm1.partialUpfront": "0.203",
            "yrTerm3.partialUpfront": "0.1545"
          },
          "ondemand": "0.291"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.105",
            "yrTerm3.allUpfront": "0.0583",
            "yrTerm1.allUpfront": "0.0881",
            "yrTerm1.partialUpfront": "0.0897",
            "yrTerm3.partialUpfront": "0.0623"
          },
          "ondemand": "0.166"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.274",
            "yrTerm3.allUpfront": "0.1924",
            "yrTerm1.allUpfront": "0.2299",
            "yrTerm1.partialUpfront": "0.235",
            "yrTerm3.partialUpfront": "0.2046"
          },
          "ondemand": "0.401"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.844",
            "yrTerm3.allUpfront": "0.6076",
            "yrTerm1.allUpfront": "0.7068",
            "yrTerm1.partialUpfront": "0.7217",
            "yrTerm3.partialUpfront": "0.6463"
          },
          "ondemand": "0.961"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.253",
            "yrTerm3.allUpfront": "0.1743",
            "yrTerm1.allUpfront": "0.2184",
            "yrTerm1.partialUpfront": "0.2221",
            "yrTerm3.partialUpfront": "0.1819"
          },
          "ondemand": "0.3"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.149",
            "yrTerm3.allUpfront": "0.0842",
            "yrTerm1.allUpfront": "0.1272",
            "yrTerm1.partialUpfront": "0.1302",
            "yrTerm3.partialUpfront": "0.0896"
          },
          "ondemand": "0.2"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.3211",
            "yrTerm3.allUpfront": "0.2536",
            "yrTerm1.allUpfront": "0.2964",
            "yrTerm1.partialUpfront": "0.299",
            "yrTerm3.partialUpfront": "0.259"
          },
          "ondemand": "0.369"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.9811",
            "yrTerm3.allUpfront": "0.9136",
            "yrTerm1.allUpfront": "0.9564",
            "yrTerm1.partialUpfront": "0.959",
            "yrTerm3.partialUpfront": "0.919"
          },
          "ondemand": "1.029"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.2441",
            "yrTerm3.allUpfront": "0.1766",
            "yrTerm1.allUpfront": "0.2194",
            "yrTerm1.partialUpfront": "0.222",
            "yrTerm3.partialUpfront": "0.182"
          },
          "ondemand": "0.292"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.1521",
            "yrTerm3.allUpfront": "0.0846",
            "yrTerm1.allUpfront": "0.1274",
            "yrTerm1.partialUpfront": "0.13",
            "yrTerm3.partialUpfront": "0.09"
          },
          "ondemand": "0.2"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.294",
            "yrTerm3.allUpfront": "0.2057",
            "yrTerm1.allUpfront": "0.2458",
            "yrTerm1.partialUpfront": "0.2512",
            "yrTerm3.partialUpfront": "0.2188"
          },
          "ondemand": "0.42"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.861",
            "yrTerm3.allUpfront": "0.6187",
            "yrTerm1.allUpfront": "0.7213",
            "yrTerm1.partialUpfront": "0.7365",
            "yrTerm3.partialUpfront": "0.6581"
          },
          "ondemand": "0.98"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.27",
            "yrTerm3.allUpfront": "0.1788",
            "yrTerm1.allUpfront": "0.2261",
            "yrTerm1.partialUpfront": "0.2312",
            "yrTerm3.partialUpfront": "0.1902"
          },
          "ondemand": "0.313"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.149",
            "yrTerm3.allUpfront": "0.0842",
            "yrTerm1.allUpfront": "0.1272",
            "yrTerm1.partialUpfront": "0.1302",
            "yrTerm3.partialUpfront": "0.0896"
          },
          "ondemand": "0.2"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.294",
            "yrTerm3.allUpfront": "0.2057",
            "yrTerm1.allUpfront": "0.2458",
            "yrTerm1.partialUpfront": "0.2512",
            "yrTerm3.partialUpfront": "0.2188"
          },
          "ondemand": "0.42"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.861",
            "yrTerm3.allUpfront": "0.6187",
            "yrTerm1.allUpfront": "0.7213",
            "yrTerm1.partialUpfront": "0.7365",
            "yrTerm3.partialUpfront": "0.6581"
          },
          "ondemand": "0.98"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.27",
            "yrTerm3.allUpfront": "0.1788",
            "yrTerm1.allUpfront": "0.2261",
            "yrTerm1.partialUpfront": "0.2312",
            "yrTerm3.partialUpfront": "0.1902"
          },
          "ondemand": "0.313"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.149",
            "yrTerm3.allUpfront": "0.0842",
            "yrTerm1.allUpfront": "0.1272",
            "yrTerm1.partialUpfront": "0.1302",
            "yrTerm3.partialUpfront": "0.0896"
          },
          "ondemand": "0.2"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.272",
            "yrTerm3.allUpfront": "0.1899",
            "yrTerm1.allUpfront": "0.2276",
            "yrTerm1.partialUpfront": "0.2319",
            "yrTerm3.partialUpfront": "0.2023"
          },
          "ondemand": "0.386"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.833",
            "yrTerm3.allUpfront": "0.5985",
            "yrTerm1.allUpfront": "0.6974",
            "yrTerm1.partialUpfront": "0.7113",
            "yrTerm3.partialUpfront": "0.637"
          },
          "ondemand": "0.946"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.238",
            "yrTerm3.allUpfront": "0.1457",
            "yrTerm1.allUpfront": "0.1992",
            "yrTerm1.partialUpfront": "0.203",
            "yrTerm3.partialUpfront": "0.1545"
          },
          "ondemand": "0.291"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.105",
            "yrTerm3.allUpfront": "0.0583",
            "yrTerm1.allUpfront": "0.0881",
            "yrTerm1.partialUpfront": "0.0897",
            "yrTerm3.partialUpfront": "0.0623"
          },
          "ondemand": "0.166"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.293",
            "yrTerm3.allUpfront": "0.2019",
            "yrTerm1.allUpfront": "0.2455",
            "yrTerm1.partialUpfront": "0.2507",
            "yrTerm3.partialUpfront": "0.2144"
          },
          "ondemand": "0.42"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.854",
            "yrTerm3.allUpfront": "0.6105",
            "yrTerm1.allUpfront": "0.7154",
            "yrTerm1.partialUpfront": "0.7301",
            "yrTerm3.partialUpfront": "0.6491"
          },
          "ondemand": "0.98"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.259",
            "yrTerm3.allUpfront": "0.1587",
            "yrTerm1.allUpfront": "0.2171",
            "yrTerm1.partialUpfront": "0.2217",
            "yrTerm3.partialUpfront": "0.1667"
          },
          "ondemand": "0.325"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.125",
            "yrTerm3.allUpfront": "0.0699",
            "yrTerm1.allUpfront": "0.1058",
            "yrTerm1.partialUpfront": "0.1074",
            "yrTerm3.partialUpfront": "0.0738"
          },
          "ondemand": "0.2"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.29",
            "yrTerm3.allUpfront": "0.2028",
            "yrTerm1.allUpfront": "0.2429",
            "yrTerm1.partialUpfront": "0.2481",
            "yrTerm3.partialUpfront": "0.2154"
          },
          "ondemand": "0.414"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.856",
            "yrTerm3.allUpfront": "0.6158",
            "yrTerm1.allUpfront": "0.7169",
            "yrTerm1.partialUpfront": "0.7317",
            "yrTerm3.partialUpfront": "0.6548"
          },
          "ondemand": "0.974"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.251",
            "yrTerm3.allUpfront": "0.1786",
            "yrTerm1.allUpfront": "0.2102",
            "yrTerm1.partialUpfront": "0.2146",
            "yrTerm3.partialUpfront": "0.1897"
          },
          "ondemand": "0.32"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.136",
            "yrTerm3.allUpfront": "0.0756",
            "yrTerm1.allUpfront": "0.1142",
            "yrTerm1.partialUpfront": "0.1167",
            "yrTerm3.partialUpfront": "0.0802"
          },
          "ondemand": "0.185"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.313",
            "yrTerm3.allUpfront": "0.2131",
            "yrTerm1.allUpfront": "0.2624",
            "yrTerm1.partialUpfront": "0.268",
            "yrTerm3.partialUpfront": "0.2264"
          },
          "ondemand": "0.508"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.874",
            "yrTerm3.allUpfront": "0.6217",
            "yrTerm1.allUpfront": "0.7323",
            "yrTerm1.partialUpfront": "0.7474",
            "yrTerm3.partialUpfront": "0.6611"
          },
          "ondemand": "1.292"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.279",
            "yrTerm3.allUpfront": "0.1808",
            "yrTerm1.allUpfront": "0.234",
            "yrTerm1.partialUpfront": "0.239",
            "yrTerm3.partialUpfront": "0.192"
          },
          "ondemand": "0.325"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.146",
            "yrTerm3.allUpfront": "0.0817",
            "yrTerm1.allUpfront": "0.1232",
            "yrTerm1.partialUpfront": "0.1259",
            "yrTerm3.partialUpfront": "0.0866"
          },
          "ondemand": "0.2"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.27",
            "yrTerm3.allUpfront": "0.189",
            "yrTerm1.allUpfront": "0.2263",
            "yrTerm1.partialUpfront": "0.231",
            "yrTerm3.partialUpfront": "0.2007"
          },
          "ondemand": "0.385"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.831",
            "yrTerm3.allUpfront": "0.5976",
            "yrTerm1.allUpfront": "0.696",
            "yrTerm1.partialUpfront": "0.7104",
            "yrTerm3.partialUpfront": "0.6354"
          },
          "ondemand": "0.945"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.251",
            "yrTerm3.allUpfront": "0.1786",
            "yrTerm1.allUpfront": "0.2102",
            "yrTerm1.partialUpfront": "0.2146",
            "yrTerm3.partialUpfront": "0.1897"
          },
          "ondemand": "0.29"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.136",
            "yrTerm3.allUpfront": "0.0756",
            "yrTerm1.allUpfront": "0.1142",
            "yrTerm1.partialUpfront": "0.1167",
            "yrTerm3.partialUpfront": "0.0802"
          },
          "ondemand": "0.185"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 10,
      "max_enis": 3
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": false,
    "storage": {
      "ssd": true,
      "devices": 1,
      "size": 32.0
    },
    "instance_type": "r3.large",
    "ECU": 6.5,
    "memory": 15.25,
    "ebs_max_bandwidth": 0
  },
  {
    "family": "Memory optimized",
    "enhanced_networking": true,
    "vCPU": 4,
    "generation": "current",
    "ebs_iops": 62.5,
    "network_performance": "Moderate",
    "ebs_throughput": 4000.0,
    "pretty_name": "R3 High-Memory Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.519",
            "yrTerm3.allUpfront": "0.3637",
            "yrTerm1.allUpfront": "0.435",
            "yrTerm1.partialUpfront": "0.4442",
            "yrTerm3.partialUpfront": "0.3865"
          },
          "ondemand": "0.738"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.213",
            "yrTerm3.allUpfront": "0.8726",
            "yrTerm1.allUpfront": "1.0158",
            "yrTerm1.partialUpfront": "1.0367",
            "yrTerm3.partialUpfront": "0.9278"
          },
          "ondemand": "1.378"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.428",
            "yrTerm3.allUpfront": "0.2632",
            "yrTerm1.allUpfront": "0.3604",
            "yrTerm1.partialUpfront": "0.3657",
            "yrTerm3.partialUpfront": "0.28"
          },
          "ondemand": "0.583"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.209",
            "yrTerm3.allUpfront": "0.1166",
            "yrTerm1.allUpfront": "0.1764",
            "yrTerm1.partialUpfront": "0.1804",
            "yrTerm3.partialUpfront": "0.1237"
          },
          "ondemand": "0.333"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.541",
            "yrTerm3.allUpfront": "0.3782",
            "yrTerm1.allUpfront": "0.4533",
            "yrTerm1.partialUpfront": "0.4623",
            "yrTerm3.partialUpfront": "0.4021"
          },
          "ondemand": "0.787"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.231",
            "yrTerm3.allUpfront": "0.8856",
            "yrTerm1.allUpfront": "1.0309",
            "yrTerm1.partialUpfront": "1.0517",
            "yrTerm3.partialUpfront": "0.9419"
          },
          "ondemand": "1.427"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.489",
            "yrTerm3.allUpfront": "0.3384",
            "yrTerm1.allUpfront": "0.4232",
            "yrTerm1.partialUpfront": "0.4318",
            "yrTerm3.partialUpfront": "0.36"
          },
          "ondemand": "0.599"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.298",
            "yrTerm3.allUpfront": "0.1677",
            "yrTerm1.allUpfront": "0.2546",
            "yrTerm1.partialUpfront": "0.2595",
            "yrTerm3.partialUpfront": "0.1782"
          },
          "ondemand": "0.399"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.5612",
            "yrTerm3.allUpfront": "0.4243",
            "yrTerm1.allUpfront": "0.5118",
            "yrTerm1.partialUpfront": "0.517",
            "yrTerm3.partialUpfront": "0.435"
          },
          "ondemand": "0.656"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.1962",
            "yrTerm3.allUpfront": "1.0593",
            "yrTerm1.allUpfront": "1.1468",
            "yrTerm1.partialUpfront": "1.152",
            "yrTerm3.partialUpfront": "1.07"
          },
          "ondemand": "1.291"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.4882",
            "yrTerm3.allUpfront": "0.3513",
            "yrTerm1.allUpfront": "0.4388",
            "yrTerm1.partialUpfront": "0.444",
            "yrTerm3.partialUpfront": "0.362"
          },
          "ondemand": "0.583"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.3042",
            "yrTerm3.allUpfront": "0.1673",
            "yrTerm1.allUpfront": "0.2548",
            "yrTerm1.partialUpfront": "0.26",
            "yrTerm3.partialUpfront": "0.178"
          },
          "ondemand": "0.399"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.579",
            "yrTerm3.allUpfront": "0.4048",
            "yrTerm1.allUpfront": "0.4847",
            "yrTerm1.partialUpfront": "0.4943",
            "yrTerm3.partialUpfront": "0.4304"
          },
          "ondemand": "0.826"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.288",
            "yrTerm3.allUpfront": "0.9278",
            "yrTerm1.allUpfront": "1.0791",
            "yrTerm1.partialUpfront": "1.1009",
            "yrTerm3.partialUpfront": "0.9868"
          },
          "ondemand": "1.466"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.523",
            "yrTerm3.allUpfront": "0.3517",
            "yrTerm1.allUpfront": "0.4385",
            "yrTerm1.partialUpfront": "0.4469",
            "yrTerm3.partialUpfront": "0.3682"
          },
          "ondemand": "0.625"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.298",
            "yrTerm3.allUpfront": "0.1677",
            "yrTerm1.allUpfront": "0.2546",
            "yrTerm1.partialUpfront": "0.2595",
            "yrTerm3.partialUpfront": "0.1782"
          },
          "ondemand": "0.399"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.579",
            "yrTerm3.allUpfront": "0.4048",
            "yrTerm1.allUpfront": "0.4847",
            "yrTerm1.partialUpfront": "0.4943",
            "yrTerm3.partialUpfront": "0.4304"
          },
          "ondemand": "0.826"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.288",
            "yrTerm3.allUpfront": "0.9278",
            "yrTerm1.allUpfront": "1.0791",
            "yrTerm1.partialUpfront": "1.1009",
            "yrTerm3.partialUpfront": "0.9868"
          },
          "ondemand": "1.466"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.523",
            "yrTerm3.allUpfront": "0.3517",
            "yrTerm1.allUpfront": "0.4385",
            "yrTerm1.partialUpfront": "0.4469",
            "yrTerm3.partialUpfront": "0.3682"
          },
          "ondemand": "0.625"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.298",
            "yrTerm3.allUpfront": "0.1677",
            "yrTerm1.allUpfront": "0.2546",
            "yrTerm1.partialUpfront": "0.2595",
            "yrTerm3.partialUpfront": "0.1782"
          },
          "ondemand": "0.399"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.519",
            "yrTerm3.allUpfront": "0.3637",
            "yrTerm1.allUpfront": "0.435",
            "yrTerm1.partialUpfront": "0.4442",
            "yrTerm3.partialUpfront": "0.3865"
          },
          "ondemand": "0.738"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.213",
            "yrTerm3.allUpfront": "0.8726",
            "yrTerm1.allUpfront": "1.0158",
            "yrTerm1.partialUpfront": "1.0367",
            "yrTerm3.partialUpfront": "0.9278"
          },
          "ondemand": "1.378"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.428",
            "yrTerm3.allUpfront": "0.2632",
            "yrTerm1.allUpfront": "0.3604",
            "yrTerm1.partialUpfront": "0.3657",
            "yrTerm3.partialUpfront": "0.28"
          },
          "ondemand": "0.583"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.209",
            "yrTerm3.allUpfront": "0.1166",
            "yrTerm1.allUpfront": "0.1764",
            "yrTerm1.partialUpfront": "0.1804",
            "yrTerm3.partialUpfront": "0.1237"
          },
          "ondemand": "0.333"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.563",
            "yrTerm3.allUpfront": "0.3878",
            "yrTerm1.allUpfront": "0.471",
            "yrTerm1.partialUpfront": "0.4806",
            "yrTerm3.partialUpfront": "0.4127"
          },
          "ondemand": "0.804"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.256",
            "yrTerm3.allUpfront": "0.8966",
            "yrTerm1.allUpfront": "1.0516",
            "yrTerm1.partialUpfront": "1.0731",
            "yrTerm3.partialUpfront": "0.954"
          },
          "ondemand": "1.444"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.471",
            "yrTerm3.allUpfront": "0.2882",
            "yrTerm1.allUpfront": "0.3957",
            "yrTerm1.partialUpfront": "0.4021",
            "yrTerm3.partialUpfront": "0.3067"
          },
          "ondemand": "0.649"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.251",
            "yrTerm3.allUpfront": "0.1399",
            "yrTerm1.allUpfront": "0.2116",
            "yrTerm1.partialUpfront": "0.2158",
            "yrTerm3.partialUpfront": "0.1486"
          },
          "ondemand": "0.399"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.552",
            "yrTerm3.allUpfront": "0.3862",
            "yrTerm1.allUpfront": "0.462",
            "yrTerm1.partialUpfront": "0.4716",
            "yrTerm3.partialUpfront": "0.4113"
          },
          "ondemand": "0.787"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.256",
            "yrTerm3.allUpfront": "0.9027",
            "yrTerm1.allUpfront": "1.0516",
            "yrTerm1.partialUpfront": "1.0733",
            "yrTerm3.partialUpfront": "0.9608"
          },
          "ondemand": "1.427"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.456",
            "yrTerm3.allUpfront": "0.329",
            "yrTerm1.allUpfront": "0.3913",
            "yrTerm1.partialUpfront": "0.3994",
            "yrTerm3.partialUpfront": "0.3447"
          },
          "ondemand": "0.633"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.272",
            "yrTerm3.allUpfront": "0.1513",
            "yrTerm1.allUpfront": "0.2283",
            "yrTerm1.partialUpfront": "0.2333",
            "yrTerm3.partialUpfront": "0.1613"
          },
          "ondemand": "0.371"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.603",
            "yrTerm3.allUpfront": "0.4101",
            "yrTerm1.allUpfront": "0.5047",
            "yrTerm1.partialUpfront": "0.5154",
            "yrTerm3.partialUpfront": "0.4367"
          },
          "ondemand": "0.967"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.296",
            "yrTerm3.allUpfront": "0.919",
            "yrTerm1.allUpfront": "1.0855",
            "yrTerm1.partialUpfront": "1.108",
            "yrTerm3.partialUpfront": "0.978"
          },
          "ondemand": "1.863"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.511",
            "yrTerm3.allUpfront": "0.3316",
            "yrTerm1.allUpfront": "0.4304",
            "yrTerm1.partialUpfront": "0.4426",
            "yrTerm3.partialUpfront": "0.3498"
          },
          "ondemand": "0.65"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.293",
            "yrTerm3.allUpfront": "0.1634",
            "yrTerm1.allUpfront": "0.2463",
            "yrTerm1.partialUpfront": "0.2519",
            "yrTerm3.partialUpfront": "0.1742"
          },
          "ondemand": "0.4"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.516",
            "yrTerm3.allUpfront": "0.3619",
            "yrTerm1.allUpfront": "0.4323",
            "yrTerm1.partialUpfront": "0.4414",
            "yrTerm3.partialUpfront": "0.3854"
          },
          "ondemand": "0.736"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.21",
            "yrTerm3.allUpfront": "0.8707",
            "yrTerm1.allUpfront": "1.013",
            "yrTerm1.partialUpfront": "1.0339",
            "yrTerm3.partialUpfront": "0.9267"
          },
          "ondemand": "1.376"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.456",
            "yrTerm3.allUpfront": "0.329",
            "yrTerm1.allUpfront": "0.3913",
            "yrTerm1.partialUpfront": "0.3994",
            "yrTerm3.partialUpfront": "0.3447"
          },
          "ondemand": "0.581"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.272",
            "yrTerm3.allUpfront": "0.1513",
            "yrTerm1.allUpfront": "0.2283",
            "yrTerm1.partialUpfront": "0.2333",
            "yrTerm3.partialUpfront": "0.1613"
          },
          "ondemand": "0.371"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 15,
      "max_enis": 4
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": true,
    "storage": {
      "ssd": true,
      "devices": 1,
      "size": 80.0
    },
    "instance_type": "r3.xlarge",
    "ECU": 13.0,
    "memory": 30.5,
    "ebs_max_bandwidth": 500.0
  },
  {
    "family": "Memory optimized",
    "enhanced_networking": true,
    "vCPU": 8,
    "generation": "current",
    "ebs_iops": 125.0,
    "network_performance": "High",
    "ebs_throughput": 8000.0,
    "pretty_name": "R3 High-Memory Double Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.069",
            "yrTerm3.allUpfront": "0.748",
            "yrTerm1.allUpfront": "0.8952",
            "yrTerm1.partialUpfront": "0.9132",
            "yrTerm3.partialUpfront": "0.796"
          },
          "ondemand": "1.52"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.412",
            "yrTerm3.allUpfront": "1.7355",
            "yrTerm1.allUpfront": "2.0204",
            "yrTerm1.partialUpfront": "2.0613",
            "yrTerm3.partialUpfront": "1.8465"
          },
          "ondemand": "2.74"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.824",
            "yrTerm3.allUpfront": "0.5076",
            "yrTerm1.allUpfront": "0.7089",
            "yrTerm1.partialUpfront": "0.7233",
            "yrTerm3.partialUpfront": "0.54"
          },
          "ondemand": "1.045"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.418",
            "yrTerm3.allUpfront": "0.2333",
            "yrTerm1.allUpfront": "0.3527",
            "yrTerm1.partialUpfront": "0.3597",
            "yrTerm3.partialUpfront": "0.2484"
          },
          "ondemand": "0.665"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.112",
            "yrTerm3.allUpfront": "0.7788",
            "yrTerm1.allUpfront": "0.9307",
            "yrTerm1.partialUpfront": "0.9494",
            "yrTerm3.partialUpfront": "0.829"
          },
          "ondemand": "1.618"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.478",
            "yrTerm3.allUpfront": "1.7821",
            "yrTerm1.allUpfront": "2.0751",
            "yrTerm1.partialUpfront": "2.1171",
            "yrTerm3.partialUpfront": "1.8964"
          },
          "ondemand": "2.872"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.965",
            "yrTerm3.allUpfront": "0.6529",
            "yrTerm1.allUpfront": "0.8159",
            "yrTerm1.partialUpfront": "0.8326",
            "yrTerm3.partialUpfront": "0.6946"
          },
          "ondemand": "1.177"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.597",
            "yrTerm3.allUpfront": "0.3354",
            "yrTerm1.allUpfront": "0.509",
            "yrTerm1.partialUpfront": "0.5191",
            "yrTerm3.partialUpfront": "0.3573"
          },
          "ondemand": "0.798"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.1212",
            "yrTerm3.allUpfront": "0.8496",
            "yrTerm1.allUpfront": "1.0226",
            "yrTerm1.partialUpfront": "1.033",
            "yrTerm3.partialUpfront": "0.871"
          },
          "ondemand": "1.312"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.3912",
            "yrTerm3.allUpfront": "2.1196",
            "yrTerm1.allUpfront": "2.2926",
            "yrTerm1.partialUpfront": "2.303",
            "yrTerm3.partialUpfront": "2.141"
          },
          "ondemand": "2.582"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.9752",
            "yrTerm3.allUpfront": "0.7036",
            "yrTerm1.allUpfront": "0.8766",
            "yrTerm1.partialUpfront": "0.887",
            "yrTerm3.partialUpfront": "0.725"
          },
          "ondemand": "1.166"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.6072",
            "yrTerm3.allUpfront": "0.3356",
            "yrTerm1.allUpfront": "0.5086",
            "yrTerm1.partialUpfront": "0.519",
            "yrTerm3.partialUpfront": "0.357"
          },
          "ondemand": "0.798"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.152",
            "yrTerm3.allUpfront": "0.8062",
            "yrTerm1.allUpfront": "0.9642",
            "yrTerm1.partialUpfront": "0.9835",
            "yrTerm3.partialUpfront": "0.8581"
          },
          "ondemand": "1.643"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.563",
            "yrTerm3.allUpfront": "1.8441",
            "yrTerm1.allUpfront": "2.1459",
            "yrTerm1.partialUpfront": "2.1894",
            "yrTerm3.partialUpfront": "1.9624"
          },
          "ondemand": "2.914"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.009",
            "yrTerm3.allUpfront": "0.6857",
            "yrTerm1.allUpfront": "0.8713",
            "yrTerm1.partialUpfront": "0.8874",
            "yrTerm3.partialUpfront": "0.725"
          },
          "ondemand": "1.25"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.597",
            "yrTerm3.allUpfront": "0.3354",
            "yrTerm1.allUpfront": "0.509",
            "yrTerm1.partialUpfront": "0.5191",
            "yrTerm3.partialUpfront": "0.3573"
          },
          "ondemand": "0.798"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.152",
            "yrTerm3.allUpfront": "0.8062",
            "yrTerm1.allUpfront": "0.9642",
            "yrTerm1.partialUpfront": "0.9835",
            "yrTerm3.partialUpfront": "0.8581"
          },
          "ondemand": "1.643"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.563",
            "yrTerm3.allUpfront": "1.8441",
            "yrTerm1.allUpfront": "2.1459",
            "yrTerm1.partialUpfront": "2.1894",
            "yrTerm3.partialUpfront": "1.9624"
          },
          "ondemand": "2.914"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.009",
            "yrTerm3.allUpfront": "0.6857",
            "yrTerm1.allUpfront": "0.8713",
            "yrTerm1.partialUpfront": "0.8874",
            "yrTerm3.partialUpfront": "0.725"
          },
          "ondemand": "1.25"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.597",
            "yrTerm3.allUpfront": "0.3354",
            "yrTerm1.allUpfront": "0.509",
            "yrTerm1.partialUpfront": "0.5191",
            "yrTerm3.partialUpfront": "0.3573"
          },
          "ondemand": "0.798"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.069",
            "yrTerm3.allUpfront": "0.748",
            "yrTerm1.allUpfront": "0.8952",
            "yrTerm1.partialUpfront": "0.9132",
            "yrTerm3.partialUpfront": "0.796"
          },
          "ondemand": "1.52"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.412",
            "yrTerm3.allUpfront": "1.7355",
            "yrTerm1.allUpfront": "2.0204",
            "yrTerm1.partialUpfront": "2.0613",
            "yrTerm3.partialUpfront": "1.8465"
          },
          "ondemand": "2.74"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.824",
            "yrTerm3.allUpfront": "0.5076",
            "yrTerm1.allUpfront": "0.7089",
            "yrTerm1.partialUpfront": "0.7233",
            "yrTerm3.partialUpfront": "0.54"
          },
          "ondemand": "1.045"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.418",
            "yrTerm3.allUpfront": "0.2333",
            "yrTerm1.allUpfront": "0.3527",
            "yrTerm1.partialUpfront": "0.3597",
            "yrTerm3.partialUpfront": "0.2484"
          },
          "ondemand": "0.665"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.154",
            "yrTerm3.allUpfront": "0.7954",
            "yrTerm1.allUpfront": "0.9663",
            "yrTerm1.partialUpfront": "0.986",
            "yrTerm3.partialUpfront": "0.8464"
          },
          "ondemand": "1.653"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.498",
            "yrTerm3.allUpfront": "1.7828",
            "yrTerm1.allUpfront": "2.0914",
            "yrTerm1.partialUpfront": "2.1341",
            "yrTerm3.partialUpfront": "1.8968"
          },
          "ondemand": "2.873"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.91",
            "yrTerm3.allUpfront": "0.5568",
            "yrTerm1.allUpfront": "0.7836",
            "yrTerm1.partialUpfront": "0.7986",
            "yrTerm3.partialUpfront": "0.5924"
          },
          "ondemand": "1.178"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.502",
            "yrTerm3.allUpfront": "0.2799",
            "yrTerm1.allUpfront": "0.4234",
            "yrTerm1.partialUpfront": "0.4315",
            "yrTerm3.partialUpfront": "0.2972"
          },
          "ondemand": "0.798"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.128",
            "yrTerm3.allUpfront": "0.7897",
            "yrTerm1.allUpfront": "0.9451",
            "yrTerm1.partialUpfront": "0.9639",
            "yrTerm3.partialUpfront": "0.8399"
          },
          "ondemand": "1.609"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.488",
            "yrTerm3.allUpfront": "1.7911",
            "yrTerm1.allUpfront": "2.0841",
            "yrTerm1.partialUpfront": "2.1263",
            "yrTerm3.partialUpfront": "1.9052"
          },
          "ondemand": "2.829"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.9",
            "yrTerm3.allUpfront": "0.6345",
            "yrTerm1.allUpfront": "0.7539",
            "yrTerm1.partialUpfront": "0.7692",
            "yrTerm3.partialUpfront": "0.675"
          },
          "ondemand": "1.134"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.543",
            "yrTerm3.allUpfront": "0.3025",
            "yrTerm1.allUpfront": "0.4566",
            "yrTerm1.partialUpfront": "0.4656",
            "yrTerm3.partialUpfront": "0.3217"
          },
          "ondemand": "0.741"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.237",
            "yrTerm3.allUpfront": "0.8419",
            "yrTerm1.allUpfront": "1.0358",
            "yrTerm1.partialUpfront": "1.0565",
            "yrTerm3.partialUpfront": "0.8954"
          },
          "ondemand": "1.997"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.58",
            "yrTerm3.allUpfront": "1.8293",
            "yrTerm1.allUpfront": "2.1611",
            "yrTerm1.partialUpfront": "2.2046",
            "yrTerm3.partialUpfront": "1.9459"
          },
          "ondemand": "3.705"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.992",
            "yrTerm3.allUpfront": "0.6397",
            "yrTerm1.allUpfront": "0.8567",
            "yrTerm1.partialUpfront": "0.8653",
            "yrTerm3.partialUpfront": "0.6806"
          },
          "ondemand": "1.18"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.585",
            "yrTerm3.allUpfront": "0.3268",
            "yrTerm1.allUpfront": "0.493",
            "yrTerm1.partialUpfront": "0.5027",
            "yrTerm3.partialUpfront": "0.3474"
          },
          "ondemand": "0.8"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.062",
            "yrTerm3.allUpfront": "0.7444",
            "yrTerm1.allUpfront": "0.8898",
            "yrTerm1.partialUpfront": "0.9076",
            "yrTerm3.partialUpfront": "0.7917"
          },
          "ondemand": "1.516"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.405",
            "yrTerm3.allUpfront": "1.7318",
            "yrTerm1.allUpfront": "2.0151",
            "yrTerm1.partialUpfront": "2.0558",
            "yrTerm3.partialUpfront": "1.8422"
          },
          "ondemand": "2.736"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.9",
            "yrTerm3.allUpfront": "0.6345",
            "yrTerm1.allUpfront": "0.7539",
            "yrTerm1.partialUpfront": "0.7692",
            "yrTerm3.partialUpfront": "0.675"
          },
          "ondemand": "1.08"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.543",
            "yrTerm3.allUpfront": "0.3025",
            "yrTerm1.allUpfront": "0.4566",
            "yrTerm1.partialUpfront": "0.4656",
            "yrTerm3.partialUpfront": "0.3217"
          },
          "ondemand": "0.741"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 15,
      "max_enis": 4
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": true,
    "storage": {
      "ssd": true,
      "devices": 1,
      "size": 160.0
    },
    "instance_type": "r3.2xlarge",
    "ECU": 26.0,
    "memory": 61.0,
    "ebs_max_bandwidth": 1000.0
  },
  {
    "family": "Memory optimized",
    "enhanced_networking": true,
    "vCPU": 16,
    "generation": "current",
    "ebs_iops": 250.0,
    "network_performance": "High",
    "ebs_throughput": 16000.0,
    "pretty_name": "R3 High-Memory Quadruple Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.617",
            "yrTerm3.allUpfront": "1.1312",
            "yrTerm1.allUpfront": "1.3539",
            "yrTerm1.partialUpfront": "1.3817",
            "yrTerm3.partialUpfront": "1.2038"
          },
          "ondemand": "2.295"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "4.041",
            "yrTerm3.allUpfront": "2.9076",
            "yrTerm1.allUpfront": "3.3842",
            "yrTerm1.partialUpfront": "3.4535",
            "yrTerm3.partialUpfront": "3.0936"
          },
          "ondemand": "4.585"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.49",
            "yrTerm3.allUpfront": "0.893",
            "yrTerm1.allUpfront": "1.2479",
            "yrTerm1.partialUpfront": "1.2734",
            "yrTerm3.partialUpfront": "0.95"
          },
          "ondemand": "1.944"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.836",
            "yrTerm3.allUpfront": "0.4666",
            "yrTerm1.allUpfront": "0.7057",
            "yrTerm1.partialUpfront": "0.7204",
            "yrTerm3.partialUpfront": "0.4967"
          },
          "ondemand": "1.33"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.739",
            "yrTerm3.allUpfront": "1.2205",
            "yrTerm1.allUpfront": "1.4563",
            "yrTerm1.partialUpfront": "1.4861",
            "yrTerm3.partialUpfront": "1.2984"
          },
          "ondemand": "2.613"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "4.105",
            "yrTerm3.allUpfront": "2.955",
            "yrTerm1.allUpfront": "3.4376",
            "yrTerm1.partialUpfront": "3.5078",
            "yrTerm3.partialUpfront": "3.1437"
          },
          "ondemand": "4.903"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.715",
            "yrTerm3.allUpfront": "1.1488",
            "yrTerm1.allUpfront": "1.4368",
            "yrTerm1.partialUpfront": "1.4661",
            "yrTerm3.partialUpfront": "1.2221"
          },
          "ondemand": "2.276"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.193",
            "yrTerm3.allUpfront": "0.6707",
            "yrTerm1.allUpfront": "1.0182",
            "yrTerm1.partialUpfront": "1.0391",
            "yrTerm3.partialUpfront": "0.7136"
          },
          "ondemand": "1.596"
        }
      },
      "sa-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "2.822",
            "yrTerm3.allUpfront": "1.7315",
            "yrTerm1.allUpfront": "2.3638",
            "yrTerm1.partialUpfront": "2.4121",
            "yrTerm3.partialUpfront": "1.842"
          },
          "ondemand": "3.635"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "5.321",
            "yrTerm3.allUpfront": "3.7393",
            "yrTerm1.allUpfront": "4.4571",
            "yrTerm1.partialUpfront": "4.5481",
            "yrTerm3.partialUpfront": "3.978"
          },
          "ondemand": "6.435"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.771",
            "yrTerm3.allUpfront": "1.6619",
            "yrTerm1.allUpfront": "2.3207",
            "yrTerm1.partialUpfront": "2.368",
            "yrTerm3.partialUpfront": "1.768"
          },
          "ondemand": "3.394"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.134",
            "yrTerm3.allUpfront": "1.1788",
            "yrTerm1.allUpfront": "1.7876",
            "yrTerm1.partialUpfront": "1.8241",
            "yrTerm3.partialUpfront": "1.254"
          },
          "ondemand": "2.799"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "2.2436",
            "yrTerm3.allUpfront": "1.6992",
            "yrTerm1.allUpfront": "2.0462",
            "yrTerm1.partialUpfront": "2.067",
            "yrTerm3.partialUpfront": "1.742"
          },
          "ondemand": "2.624"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "4.7836",
            "yrTerm3.allUpfront": "4.2392",
            "yrTerm1.allUpfront": "4.5862",
            "yrTerm1.partialUpfront": "4.607",
            "yrTerm3.partialUpfront": "4.282"
          },
          "ondemand": "5.164"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.9516",
            "yrTerm3.allUpfront": "1.4072",
            "yrTerm1.allUpfront": "1.7542",
            "yrTerm1.partialUpfront": "1.775",
            "yrTerm3.partialUpfront": "1.45"
          },
          "ondemand": "2.332"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.2156",
            "yrTerm3.allUpfront": "0.6712",
            "yrTerm1.allUpfront": "1.0182",
            "yrTerm1.partialUpfront": "1.039",
            "yrTerm3.partialUpfront": "0.714"
          },
          "ondemand": "1.596"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.949",
            "yrTerm3.allUpfront": "1.3664",
            "yrTerm1.allUpfront": "1.6321",
            "yrTerm1.partialUpfront": "1.6655",
            "yrTerm3.partialUpfront": "1.4536"
          },
          "ondemand": "2.783"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "4.463",
            "yrTerm3.allUpfront": "3.2121",
            "yrTerm1.allUpfront": "3.7372",
            "yrTerm1.partialUpfront": "3.8137",
            "yrTerm3.partialUpfront": "3.4172"
          },
          "ondemand": "5.073"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.832",
            "yrTerm3.allUpfront": "1.2058",
            "yrTerm1.allUpfront": "1.5341",
            "yrTerm1.partialUpfront": "1.5655",
            "yrTerm3.partialUpfront": "1.2828"
          },
          "ondemand": "2.363"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.193",
            "yrTerm3.allUpfront": "0.6707",
            "yrTerm1.allUpfront": "1.0182",
            "yrTerm1.partialUpfront": "1.0391",
            "yrTerm3.partialUpfront": "0.7136"
          },
          "ondemand": "1.596"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.949",
            "yrTerm3.allUpfront": "1.3664",
            "yrTerm1.allUpfront": "1.6321",
            "yrTerm1.partialUpfront": "1.6655",
            "yrTerm3.partialUpfront": "1.4536"
          },
          "ondemand": "2.783"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "4.463",
            "yrTerm3.allUpfront": "3.2121",
            "yrTerm1.allUpfront": "3.7372",
            "yrTerm1.partialUpfront": "3.8137",
            "yrTerm3.partialUpfront": "3.4172"
          },
          "ondemand": "5.073"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.832",
            "yrTerm3.allUpfront": "1.2058",
            "yrTerm1.allUpfront": "1.5341",
            "yrTerm1.partialUpfront": "1.5655",
            "yrTerm3.partialUpfront": "1.2828"
          },
          "ondemand": "2.363"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.193",
            "yrTerm3.allUpfront": "0.6707",
            "yrTerm1.allUpfront": "1.0182",
            "yrTerm1.partialUpfront": "1.0391",
            "yrTerm3.partialUpfront": "0.7136"
          },
          "ondemand": "1.596"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.617",
            "yrTerm3.allUpfront": "1.1312",
            "yrTerm1.allUpfront": "1.3539",
            "yrTerm1.partialUpfront": "1.3817",
            "yrTerm3.partialUpfront": "1.2038"
          },
          "ondemand": "2.295"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "4.041",
            "yrTerm3.allUpfront": "2.9076",
            "yrTerm1.allUpfront": "3.3842",
            "yrTerm1.partialUpfront": "3.4535",
            "yrTerm3.partialUpfront": "3.0936"
          },
          "ondemand": "4.585"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.49",
            "yrTerm3.allUpfront": "0.893",
            "yrTerm1.allUpfront": "1.2479",
            "yrTerm1.partialUpfront": "1.2734",
            "yrTerm3.partialUpfront": "0.95"
          },
          "ondemand": "1.944"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.836",
            "yrTerm3.allUpfront": "0.4666",
            "yrTerm1.allUpfront": "0.7057",
            "yrTerm1.partialUpfront": "0.7204",
            "yrTerm3.partialUpfront": "0.4967"
          },
          "ondemand": "1.33"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.786",
            "yrTerm3.allUpfront": "1.2249",
            "yrTerm1.allUpfront": "1.4959",
            "yrTerm1.partialUpfront": "1.5262",
            "yrTerm3.partialUpfront": "1.3036"
          },
          "ondemand": "2.561"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "4.21",
            "yrTerm3.allUpfront": "3.0014",
            "yrTerm1.allUpfront": "3.5263",
            "yrTerm1.partialUpfront": "3.598",
            "yrTerm3.partialUpfront": "3.1934"
          },
          "ondemand": "4.851"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.668",
            "yrTerm3.allUpfront": "0.9914",
            "yrTerm1.allUpfront": "1.3974",
            "yrTerm1.partialUpfront": "1.4259",
            "yrTerm3.partialUpfront": "1.0547"
          },
          "ondemand": "2.224"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.003",
            "yrTerm3.allUpfront": "0.5599",
            "yrTerm1.allUpfront": "0.8469",
            "yrTerm1.partialUpfront": "0.8631",
            "yrTerm3.partialUpfront": "0.5954"
          },
          "ondemand": "1.596"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.721",
            "yrTerm3.allUpfront": "1.2062",
            "yrTerm1.allUpfront": "1.4409",
            "yrTerm1.partialUpfront": "1.4705",
            "yrTerm3.partialUpfront": "1.2828"
          },
          "ondemand": "2.454"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "4.175",
            "yrTerm3.allUpfront": "3.0054",
            "yrTerm1.allUpfront": "3.4966",
            "yrTerm1.partialUpfront": "3.5682",
            "yrTerm3.partialUpfront": "3.1969"
          },
          "ondemand": "4.744"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.584",
            "yrTerm3.allUpfront": "1.1163",
            "yrTerm1.allUpfront": "1.3268",
            "yrTerm1.partialUpfront": "1.3539",
            "yrTerm3.partialUpfront": "1.1875"
          },
          "ondemand": "2.111"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.087",
            "yrTerm3.allUpfront": "0.6051",
            "yrTerm1.allUpfront": "0.9132",
            "yrTerm1.partialUpfront": "0.9321",
            "yrTerm3.partialUpfront": "0.6434"
          },
          "ondemand": "1.482"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.952",
            "yrTerm3.allUpfront": "1.3179",
            "yrTerm1.allUpfront": "1.6348",
            "yrTerm1.partialUpfront": "1.6684",
            "yrTerm3.partialUpfront": "1.4017"
          },
          "ondemand": "2.951"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "4.375",
            "yrTerm3.allUpfront": "3.0943",
            "yrTerm1.allUpfront": "3.6643",
            "yrTerm1.partialUpfront": "3.7393",
            "yrTerm3.partialUpfront": "3.2914"
          },
          "ondemand": "6.157"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.843",
            "yrTerm3.allUpfront": "1.1394",
            "yrTerm1.allUpfront": "1.5437",
            "yrTerm1.partialUpfront": "1.5752",
            "yrTerm3.partialUpfront": "1.2121"
          },
          "ondemand": "2.228"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.17",
            "yrTerm3.allUpfront": "0.6535",
            "yrTerm1.allUpfront": "0.9862",
            "yrTerm1.partialUpfront": "1.0067",
            "yrTerm3.partialUpfront": "0.6948"
          },
          "ondemand": "1.6"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.604",
            "yrTerm3.allUpfront": "1.1239",
            "yrTerm1.allUpfront": "1.3429",
            "yrTerm1.partialUpfront": "1.3705",
            "yrTerm3.partialUpfront": "1.1952"
          },
          "ondemand": "2.287"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "4.028",
            "yrTerm3.allUpfront": "2.9003",
            "yrTerm1.allUpfront": "3.3733",
            "yrTerm1.partialUpfront": "3.4423",
            "yrTerm3.partialUpfront": "3.0851"
          },
          "ondemand": "4.577"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.584",
            "yrTerm3.allUpfront": "1.1163",
            "yrTerm1.allUpfront": "1.3268",
            "yrTerm1.partialUpfront": "1.3539",
            "yrTerm3.partialUpfront": "1.1875"
          },
          "ondemand": "1.944"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.087",
            "yrTerm3.allUpfront": "0.6051",
            "yrTerm1.allUpfront": "0.9132",
            "yrTerm1.partialUpfront": "0.9321",
            "yrTerm3.partialUpfront": "0.6434"
          },
          "ondemand": "1.482"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 30,
      "max_enis": 8
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": true,
    "storage": {
      "ssd": true,
      "devices": 1,
      "size": 320.0
    },
    "instance_type": "r3.4xlarge",
    "ECU": 52.0,
    "memory": 122.0,
    "ebs_max_bandwidth": 2000.0
  },
  {
    "family": "Memory optimized",
    "enhanced_networking": true,
    "vCPU": 32,
    "generation": "current",
    "ebs_iops": 0,
    "network_performance": "10 Gigabit",
    "ebs_throughput": 0,
    "pretty_name": "R3 High-Memory Eight Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "2.717",
            "yrTerm3.allUpfront": "1.9033",
            "yrTerm1.allUpfront": "2.2749",
            "yrTerm1.partialUpfront": "2.3217",
            "yrTerm3.partialUpfront": "2.0246"
          },
          "ondemand": "3.855"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "7.594",
            "yrTerm3.allUpfront": "5.4648",
            "yrTerm1.allUpfront": "6.3603",
            "yrTerm1.partialUpfront": "6.4905",
            "yrTerm3.partialUpfront": "5.8135"
          },
          "ondemand": "8.615"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.989",
            "yrTerm3.allUpfront": "1.1562",
            "yrTerm1.allUpfront": "1.666",
            "yrTerm1.partialUpfront": "1.7",
            "yrTerm3.partialUpfront": "1.23"
          },
          "ondemand": "3.5"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.672",
            "yrTerm3.allUpfront": "0.9331",
            "yrTerm1.allUpfront": "1.4114",
            "yrTerm1.partialUpfront": "1.4407",
            "yrTerm3.partialUpfront": "0.9925"
          },
          "ondemand": "2.66"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.046",
            "yrTerm3.allUpfront": "2.1376",
            "yrTerm1.allUpfront": "2.55",
            "yrTerm1.partialUpfront": "2.6024",
            "yrTerm3.partialUpfront": "2.274"
          },
          "ondemand": "4.577"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "7.817",
            "yrTerm3.allUpfront": "5.6306",
            "yrTerm1.allUpfront": "6.5462",
            "yrTerm1.partialUpfront": "6.6802",
            "yrTerm3.partialUpfront": "5.99"
          },
          "ondemand": "9.337"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.674",
            "yrTerm3.allUpfront": "1.4863",
            "yrTerm1.allUpfront": "2.2401",
            "yrTerm1.partialUpfront": "2.2857",
            "yrTerm3.partialUpfront": "1.5812"
          },
          "ondemand": "4.25"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.386",
            "yrTerm3.allUpfront": "1.3415",
            "yrTerm1.allUpfront": "2.0362",
            "yrTerm1.partialUpfront": "2.0782",
            "yrTerm3.partialUpfront": "1.4272"
          },
          "ondemand": "3.192"
        }
      },
      "sa-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "5.641",
            "yrTerm3.allUpfront": "3.4611",
            "yrTerm1.allUpfront": "4.7256",
            "yrTerm1.partialUpfront": "4.822",
            "yrTerm3.partialUpfront": "3.682"
          },
          "ondemand": "7.758"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "10.64",
            "yrTerm3.allUpfront": "7.4768",
            "yrTerm1.allUpfront": "8.9121",
            "yrTerm1.partialUpfront": "9.094",
            "yrTerm3.partialUpfront": "7.954"
          },
          "ondemand": "15.336"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "5.541",
            "yrTerm3.allUpfront": "3.3239",
            "yrTerm1.allUpfront": "4.6413",
            "yrTerm1.partialUpfront": "4.736",
            "yrTerm3.partialUpfront": "3.536"
          },
          "ondemand": "6.788"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "4.268",
            "yrTerm3.allUpfront": "2.3575",
            "yrTerm1.allUpfront": "3.5751",
            "yrTerm1.partialUpfront": "3.648",
            "yrTerm3.partialUpfront": "2.508"
          },
          "ondemand": "5.597"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "4.4873",
            "yrTerm3.allUpfront": "3.3974",
            "yrTerm1.allUpfront": "4.0924",
            "yrTerm1.partialUpfront": "4.134",
            "yrTerm3.partialUpfront": "3.483"
          },
          "ondemand": "5.248"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "9.5673",
            "yrTerm3.allUpfront": "8.4774",
            "yrTerm1.allUpfront": "9.1724",
            "yrTerm1.partialUpfront": "9.214",
            "yrTerm3.partialUpfront": "8.563"
          },
          "ondemand": "10.328"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "3.9033",
            "yrTerm3.allUpfront": "2.8134",
            "yrTerm1.allUpfront": "3.5084",
            "yrTerm1.partialUpfront": "3.55",
            "yrTerm3.partialUpfront": "2.899"
          },
          "ondemand": "4.664"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.4313",
            "yrTerm3.allUpfront": "1.3414",
            "yrTerm1.allUpfront": "2.0364",
            "yrTerm1.partialUpfront": "2.078",
            "yrTerm3.partialUpfront": "1.427"
          },
          "ondemand": "3.192"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.451",
            "yrTerm3.allUpfront": "2.4193",
            "yrTerm1.allUpfront": "2.8888",
            "yrTerm1.partialUpfront": "2.9481",
            "yrTerm3.partialUpfront": "2.5737"
          },
          "ondemand": "4.927"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "8.522",
            "yrTerm3.allUpfront": "6.1361",
            "yrTerm1.allUpfront": "7.1361",
            "yrTerm1.partialUpfront": "7.2821",
            "yrTerm3.partialUpfront": "6.5278"
          },
          "ondemand": "9.687"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.925",
            "yrTerm3.allUpfront": "1.561",
            "yrTerm1.allUpfront": "2.45",
            "yrTerm1.partialUpfront": "2.5",
            "yrTerm3.partialUpfront": "1.6606"
          },
          "ondemand": "4.6"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.386",
            "yrTerm3.allUpfront": "1.3415",
            "yrTerm1.allUpfront": "2.0362",
            "yrTerm1.partialUpfront": "2.0782",
            "yrTerm3.partialUpfront": "1.4272"
          },
          "ondemand": "3.192"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.451",
            "yrTerm3.allUpfront": "2.4193",
            "yrTerm1.allUpfront": "2.8888",
            "yrTerm1.partialUpfront": "2.9481",
            "yrTerm3.partialUpfront": "2.5737"
          },
          "ondemand": "4.927"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "8.522",
            "yrTerm3.allUpfront": "6.1361",
            "yrTerm1.allUpfront": "7.1361",
            "yrTerm1.partialUpfront": "7.2821",
            "yrTerm3.partialUpfront": "6.5278"
          },
          "ondemand": "9.687"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.925",
            "yrTerm3.allUpfront": "1.561",
            "yrTerm1.allUpfront": "2.45",
            "yrTerm1.partialUpfront": "2.5",
            "yrTerm3.partialUpfront": "1.6606"
          },
          "ondemand": "4.6"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.386",
            "yrTerm3.allUpfront": "1.3415",
            "yrTerm1.allUpfront": "2.0362",
            "yrTerm1.partialUpfront": "2.0782",
            "yrTerm3.partialUpfront": "1.4272"
          },
          "ondemand": "3.192"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "2.717",
            "yrTerm3.allUpfront": "1.9033",
            "yrTerm1.allUpfront": "2.2749",
            "yrTerm1.partialUpfront": "2.3217",
            "yrTerm3.partialUpfront": "2.0246"
          },
          "ondemand": "3.855"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "7.594",
            "yrTerm3.allUpfront": "5.4648",
            "yrTerm1.allUpfront": "6.3603",
            "yrTerm1.partialUpfront": "6.4905",
            "yrTerm3.partialUpfront": "5.8135"
          },
          "ondemand": "8.615"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.989",
            "yrTerm3.allUpfront": "1.1562",
            "yrTerm1.allUpfront": "1.666",
            "yrTerm1.partialUpfront": "1.7",
            "yrTerm3.partialUpfront": "1.23"
          },
          "ondemand": "3.5"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.672",
            "yrTerm3.allUpfront": "0.9331",
            "yrTerm1.allUpfront": "1.4114",
            "yrTerm1.partialUpfront": "1.4407",
            "yrTerm3.partialUpfront": "0.9925"
          },
          "ondemand": "2.66"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.054",
            "yrTerm3.allUpfront": "2.0907",
            "yrTerm1.allUpfront": "2.558",
            "yrTerm1.partialUpfront": "2.61",
            "yrTerm3.partialUpfront": "2.2241"
          },
          "ondemand": "4.387"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "7.932",
            "yrTerm3.allUpfront": "5.6522",
            "yrTerm1.allUpfront": "6.6435",
            "yrTerm1.partialUpfront": "6.7788",
            "yrTerm3.partialUpfront": "6.013"
          },
          "ondemand": "9.147"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.345",
            "yrTerm3.allUpfront": "1.353",
            "yrTerm1.allUpfront": "1.964",
            "yrTerm1.partialUpfront": "2.0042",
            "yrTerm3.partialUpfront": "1.4394"
          },
          "ondemand": "4.06"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.006",
            "yrTerm3.allUpfront": "1.1197",
            "yrTerm1.allUpfront": "1.6937",
            "yrTerm1.partialUpfront": "1.7275",
            "yrTerm3.partialUpfront": "1.1909"
          },
          "ondemand": "3.192"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "2.901",
            "yrTerm3.allUpfront": "2.0357",
            "yrTerm1.allUpfront": "2.4293",
            "yrTerm1.partialUpfront": "2.4793",
            "yrTerm3.partialUpfront": "2.1659"
          },
          "ondemand": "4.139"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "7.832",
            "yrTerm3.allUpfront": "5.6389",
            "yrTerm1.allUpfront": "6.5592",
            "yrTerm1.partialUpfront": "6.6935",
            "yrTerm3.partialUpfront": "5.999"
          },
          "ondemand": "8.899"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.518",
            "yrTerm3.allUpfront": "1.4452",
            "yrTerm1.allUpfront": "2.109",
            "yrTerm1.partialUpfront": "2.1521",
            "yrTerm3.partialUpfront": "1.5375"
          },
          "ondemand": "3.8"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.174",
            "yrTerm3.allUpfront": "1.2102",
            "yrTerm1.allUpfront": "1.8266",
            "yrTerm1.partialUpfront": "1.8643",
            "yrTerm3.partialUpfront": "1.2877"
          },
          "ondemand": "2.964"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.387",
            "yrTerm3.allUpfront": "2.2777",
            "yrTerm1.allUpfront": "2.8361",
            "yrTerm1.partialUpfront": "2.8943",
            "yrTerm3.partialUpfront": "2.4233"
          },
          "ondemand": "4.874"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "8.264",
            "yrTerm3.allUpfront": "5.8392",
            "yrTerm1.allUpfront": "6.9216",
            "yrTerm1.partialUpfront": "7.0631",
            "yrTerm3.partialUpfront": "6.2121"
          },
          "ondemand": "11.538"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.694",
            "yrTerm3.allUpfront": "1.5494",
            "yrTerm1.allUpfront": "2.2567",
            "yrTerm1.partialUpfront": "2.3028",
            "yrTerm3.partialUpfront": "1.6483"
          },
          "ondemand": "4.069"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.341",
            "yrTerm3.allUpfront": "1.307",
            "yrTerm1.allUpfront": "1.9726",
            "yrTerm1.partialUpfront": "2.0132",
            "yrTerm3.partialUpfront": "1.3907"
          },
          "ondemand": "3.201"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "2.691",
            "yrTerm3.allUpfront": "1.8887",
            "yrTerm1.allUpfront": "2.2531",
            "yrTerm1.partialUpfront": "2.2994",
            "yrTerm3.partialUpfront": "2.0095"
          },
          "ondemand": "3.839"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "7.568",
            "yrTerm3.allUpfront": "5.4502",
            "yrTerm1.allUpfront": "6.3385",
            "yrTerm1.partialUpfront": "6.4682",
            "yrTerm3.partialUpfront": "5.7983"
          },
          "ondemand": "8.599"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.518",
            "yrTerm3.allUpfront": "1.4452",
            "yrTerm1.allUpfront": "2.109",
            "yrTerm1.partialUpfront": "2.1521",
            "yrTerm3.partialUpfront": "1.5375"
          },
          "ondemand": "3.5"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.174",
            "yrTerm3.allUpfront": "1.2102",
            "yrTerm1.allUpfront": "1.8266",
            "yrTerm1.partialUpfront": "1.8643",
            "yrTerm3.partialUpfront": "1.2877"
          },
          "ondemand": "2.964"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 30,
      "max_enis": 8
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": false,
    "storage": {
      "ssd": true,
      "devices": 2,
      "size": 320.0
    },
    "instance_type": "r3.8xlarge",
    "ECU": 104.0,
    "memory": 244.0,
    "ebs_max_bandwidth": 0
  },
  {
    "family": "Storage optimized",
    "enhanced_networking": true,
    "vCPU": 4,
    "generation": "current",
    "ebs_iops": 62.5,
    "network_performance": "Moderate",
    "ebs_throughput": 4000.0,
    "pretty_name": "I2 Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.578",
            "yrTerm3.allUpfront": "0.3339",
            "yrTerm1.allUpfront": "0.4839",
            "yrTerm1.partialUpfront": "0.4944",
            "yrTerm3.partialUpfront": "0.3552"
          },
          "ondemand": "0.993"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.768",
            "yrTerm3.allUpfront": "0.487",
            "yrTerm1.allUpfront": "0.6433",
            "yrTerm1.partialUpfront": "0.6564",
            "yrTerm3.partialUpfront": "0.5181"
          },
          "ondemand": "1.23"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.565",
            "yrTerm3.allUpfront": "0.3245",
            "yrTerm1.allUpfront": "0.4731",
            "yrTerm1.partialUpfront": "0.4828",
            "yrTerm3.partialUpfront": "0.3453"
          },
          "ondemand": "0.973"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.424",
            "yrTerm3.allUpfront": "0.2117",
            "yrTerm1.allUpfront": "0.3555",
            "yrTerm1.partialUpfront": "0.3628",
            "yrTerm3.partialUpfront": "0.2253"
          },
          "ondemand": "0.853"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.848",
            "yrTerm3.allUpfront": "0.4895",
            "yrTerm1.allUpfront": "0.7103",
            "yrTerm1.partialUpfront": "0.7252",
            "yrTerm3.partialUpfront": "0.5207"
          },
          "ondemand": "1.23"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.126",
            "yrTerm3.allUpfront": "0.7142",
            "yrTerm1.allUpfront": "0.9436",
            "yrTerm1.partialUpfront": "0.9628",
            "yrTerm3.partialUpfront": "0.7598"
          },
          "ondemand": "2.026"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.829",
            "yrTerm3.allUpfront": "0.4756",
            "yrTerm1.allUpfront": "0.6946",
            "yrTerm1.partialUpfront": "0.7088",
            "yrTerm3.partialUpfront": "0.5059"
          },
          "ondemand": "1.113"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.623",
            "yrTerm3.allUpfront": "0.3101",
            "yrTerm1.allUpfront": "0.522",
            "yrTerm1.partialUpfront": "0.5328",
            "yrTerm3.partialUpfront": "0.3299"
          },
          "ondemand": "1.001"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.8804",
            "yrTerm3.allUpfront": "0.5671",
            "yrTerm1.allUpfront": "0.7791",
            "yrTerm1.partialUpfront": "0.7898",
            "yrTerm3.partialUpfront": "0.5869"
          },
          "ondemand": "1.258"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.5154",
            "yrTerm3.allUpfront": "1.2021",
            "yrTerm1.allUpfront": "1.4141",
            "yrTerm1.partialUpfront": "1.4248",
            "yrTerm3.partialUpfront": "1.2219"
          },
          "ondemand": "1.893"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.8074",
            "yrTerm3.allUpfront": "0.4941",
            "yrTerm1.allUpfront": "0.7061",
            "yrTerm1.partialUpfront": "0.7168",
            "yrTerm3.partialUpfront": "0.5139"
          },
          "ondemand": "1.185"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.6234",
            "yrTerm3.allUpfront": "0.3101",
            "yrTerm1.allUpfront": "0.5221",
            "yrTerm1.partialUpfront": "0.5328",
            "yrTerm3.partialUpfront": "0.3299"
          },
          "ondemand": "1.001"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.832",
            "yrTerm3.allUpfront": "0.4811",
            "yrTerm1.allUpfront": "0.697",
            "yrTerm1.partialUpfront": "0.7116",
            "yrTerm3.partialUpfront": "0.5118"
          },
          "ondemand": "1.292"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.106",
            "yrTerm3.allUpfront": "0.7015",
            "yrTerm1.allUpfront": "0.926",
            "yrTerm1.partialUpfront": "0.945",
            "yrTerm3.partialUpfront": "0.7462"
          },
          "ondemand": "2.127"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.813",
            "yrTerm3.allUpfront": "0.4673",
            "yrTerm1.allUpfront": "0.6813",
            "yrTerm1.partialUpfront": "0.6952",
            "yrTerm3.partialUpfront": "0.4972"
          },
          "ondemand": "1.169"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.61",
            "yrTerm3.allUpfront": "0.3047",
            "yrTerm1.allUpfront": "0.5116",
            "yrTerm1.partialUpfront": "0.5222",
            "yrTerm3.partialUpfront": "0.3242"
          },
          "ondemand": "1.018"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.832",
            "yrTerm3.allUpfront": "0.4811",
            "yrTerm1.allUpfront": "0.697",
            "yrTerm1.partialUpfront": "0.7116",
            "yrTerm3.partialUpfront": "0.5118"
          },
          "ondemand": "1.292"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.106",
            "yrTerm3.allUpfront": "0.7015",
            "yrTerm1.allUpfront": "0.926",
            "yrTerm1.partialUpfront": "0.945",
            "yrTerm3.partialUpfront": "0.7462"
          },
          "ondemand": "2.127"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.813",
            "yrTerm3.allUpfront": "0.4673",
            "yrTerm1.allUpfront": "0.6813",
            "yrTerm1.partialUpfront": "0.6952",
            "yrTerm3.partialUpfront": "0.4972"
          },
          "ondemand": "1.169"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.61",
            "yrTerm3.allUpfront": "0.3047",
            "yrTerm1.allUpfront": "0.5116",
            "yrTerm1.partialUpfront": "0.5222",
            "yrTerm3.partialUpfront": "0.3242"
          },
          "ondemand": "1.018"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.578",
            "yrTerm3.allUpfront": "0.3339",
            "yrTerm1.allUpfront": "0.4839",
            "yrTerm1.partialUpfront": "0.4944",
            "yrTerm3.partialUpfront": "0.3552"
          },
          "ondemand": "0.993"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.768",
            "yrTerm3.allUpfront": "0.487",
            "yrTerm1.allUpfront": "0.6433",
            "yrTerm1.partialUpfront": "0.6564",
            "yrTerm3.partialUpfront": "0.5181"
          },
          "ondemand": "1.23"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.565",
            "yrTerm3.allUpfront": "0.3245",
            "yrTerm1.allUpfront": "0.4731",
            "yrTerm1.partialUpfront": "0.4828",
            "yrTerm3.partialUpfront": "0.3453"
          },
          "ondemand": "0.973"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.424",
            "yrTerm3.allUpfront": "0.2117",
            "yrTerm1.allUpfront": "0.3555",
            "yrTerm1.partialUpfront": "0.3628",
            "yrTerm3.partialUpfront": "0.2253"
          },
          "ondemand": "0.853"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.664",
            "yrTerm3.allUpfront": "0.377",
            "yrTerm1.allUpfront": "0.5561",
            "yrTerm1.partialUpfront": "0.5674",
            "yrTerm3.partialUpfront": "0.4011"
          },
          "ondemand": "1.164"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.854",
            "yrTerm3.allUpfront": "0.5296",
            "yrTerm1.allUpfront": "0.7148",
            "yrTerm1.partialUpfront": "0.7295",
            "yrTerm3.partialUpfront": "0.5634"
          },
          "ondemand": "1.401"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.65",
            "yrTerm3.allUpfront": "0.3672",
            "yrTerm1.allUpfront": "0.5444",
            "yrTerm1.partialUpfront": "0.5556",
            "yrTerm3.partialUpfront": "0.3906"
          },
          "ondemand": "1.144"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.509",
            "yrTerm3.allUpfront": "0.2541",
            "yrTerm1.allUpfront": "0.4266",
            "yrTerm1.partialUpfront": "0.4353",
            "yrTerm3.partialUpfront": "0.2701"
          },
          "ondemand": "1.023"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.578",
            "yrTerm3.allUpfront": "0.3339",
            "yrTerm1.allUpfront": "0.4839",
            "yrTerm1.partialUpfront": "0.4944",
            "yrTerm3.partialUpfront": "0.3552"
          },
          "ondemand": "1.078"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.96",
            "yrTerm3.allUpfront": "0.6083",
            "yrTerm1.allUpfront": "0.804",
            "yrTerm1.partialUpfront": "0.8204",
            "yrTerm3.partialUpfront": "0.6471"
          },
          "ondemand": "1.315"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.706",
            "yrTerm3.allUpfront": "0.4054",
            "yrTerm1.allUpfront": "0.5917",
            "yrTerm1.partialUpfront": "0.6038",
            "yrTerm3.partialUpfront": "0.4313"
          },
          "ondemand": "1.058"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.515",
            "yrTerm3.allUpfront": "0.2728",
            "yrTerm1.allUpfront": "0.4318",
            "yrTerm1.partialUpfront": "0.4408",
            "yrTerm3.partialUpfront": "0.2903"
          },
          "ondemand": "0.938"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.711",
            "yrTerm3.allUpfront": "0.417",
            "yrTerm1.allUpfront": "0.5952",
            "yrTerm1.partialUpfront": "0.6073",
            "yrTerm3.partialUpfront": "0.4436"
          },
          "ondemand": "1.58"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.559",
            "yrTerm3.allUpfront": "1.0507",
            "yrTerm1.allUpfront": "1.3061",
            "yrTerm1.partialUpfront": "1.3326",
            "yrTerm3.partialUpfront": "1.1178"
          },
          "ondemand": "2.476"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.709",
            "yrTerm3.allUpfront": "0.4079",
            "yrTerm1.allUpfront": "0.5936",
            "yrTerm1.partialUpfront": "0.6057",
            "yrTerm3.partialUpfront": "0.4339"
          },
          "ondemand": "1.133"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.556",
            "yrTerm3.allUpfront": "0.2946",
            "yrTerm1.allUpfront": "0.4663",
            "yrTerm1.partialUpfront": "0.4759",
            "yrTerm3.partialUpfront": "0.3135"
          },
          "ondemand": "1.013"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "0.578",
            "yrTerm3.allUpfront": "0.3339",
            "yrTerm1.allUpfront": "0.4839",
            "yrTerm1.partialUpfront": "0.4944",
            "yrTerm3.partialUpfront": "0.3552"
          },
          "ondemand": "0.993"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "0.96",
            "yrTerm3.allUpfront": "0.6083",
            "yrTerm1.allUpfront": "0.804",
            "yrTerm1.partialUpfront": "0.8204",
            "yrTerm3.partialUpfront": "0.6471"
          },
          "ondemand": "1.23"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.706",
            "yrTerm3.allUpfront": "0.4054",
            "yrTerm1.allUpfront": "0.5917",
            "yrTerm1.partialUpfront": "0.6038",
            "yrTerm3.partialUpfront": "0.4313"
          },
          "ondemand": "0.973"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.515",
            "yrTerm3.allUpfront": "0.2728",
            "yrTerm1.allUpfront": "0.4318",
            "yrTerm1.partialUpfront": "0.4408",
            "yrTerm3.partialUpfront": "0.2903"
          },
          "ondemand": "0.938"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 15,
      "max_enis": 4
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": true,
    "storage": {
      "ssd": true,
      "devices": 1,
      "size": 800.0
    },
    "instance_type": "i2.xlarge",
    "ECU": 14.0,
    "memory": 30.5,
    "ebs_max_bandwidth": 500.0
  },
  {
    "family": "Storage optimized",
    "enhanced_networking": true,
    "vCPU": 8,
    "generation": "current",
    "ebs_iops": 125.0,
    "network_performance": "High",
    "ebs_throughput": 8000.0,
    "pretty_name": "I2 Double Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.155",
            "yrTerm3.allUpfront": "0.6688",
            "yrTerm1.allUpfront": "0.9678",
            "yrTerm1.partialUpfront": "0.988",
            "yrTerm3.partialUpfront": "0.7115"
          },
          "ondemand": "1.986"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.536",
            "yrTerm3.allUpfront": "0.974",
            "yrTerm1.allUpfront": "1.2868",
            "yrTerm1.partialUpfront": "1.313",
            "yrTerm3.partialUpfront": "1.0362"
          },
          "ondemand": "2.459"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.131",
            "yrTerm3.allUpfront": "0.6481",
            "yrTerm1.allUpfront": "0.9471",
            "yrTerm1.partialUpfront": "0.9665",
            "yrTerm3.partialUpfront": "0.6895"
          },
          "ondemand": "1.946"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.848",
            "yrTerm3.allUpfront": "0.4225",
            "yrTerm1.allUpfront": "0.712",
            "yrTerm1.partialUpfront": "0.7265",
            "yrTerm3.partialUpfront": "0.4495"
          },
          "ondemand": "1.705"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.695",
            "yrTerm3.allUpfront": "0.9807",
            "yrTerm1.allUpfront": "1.4195",
            "yrTerm1.partialUpfront": "1.4486",
            "yrTerm3.partialUpfront": "1.0434"
          },
          "ondemand": "2.462"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.253",
            "yrTerm3.allUpfront": "1.4283",
            "yrTerm1.allUpfront": "1.8874",
            "yrTerm1.partialUpfront": "1.926",
            "yrTerm3.partialUpfront": "1.5195"
          },
          "ondemand": "4.051"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.66",
            "yrTerm3.allUpfront": "0.9511",
            "yrTerm1.allUpfront": "1.3901",
            "yrTerm1.partialUpfront": "1.4185",
            "yrTerm3.partialUpfront": "1.0118"
          },
          "ondemand": "2.226"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.246",
            "yrTerm3.allUpfront": "0.6193",
            "yrTerm1.allUpfront": "1.0441",
            "yrTerm1.partialUpfront": "1.0655",
            "yrTerm3.partialUpfront": "0.6588"
          },
          "ondemand": "2.001"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.7606",
            "yrTerm3.allUpfront": "1.1333",
            "yrTerm1.allUpfront": "1.5582",
            "yrTerm1.partialUpfront": "1.5795",
            "yrTerm3.partialUpfront": "1.1728"
          },
          "ondemand": "2.515"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.0306",
            "yrTerm3.allUpfront": "2.4033",
            "yrTerm1.allUpfront": "2.8282",
            "yrTerm1.partialUpfront": "2.8495",
            "yrTerm3.partialUpfront": "2.4428"
          },
          "ondemand": "3.785"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.6146",
            "yrTerm3.allUpfront": "0.9873",
            "yrTerm1.allUpfront": "1.4122",
            "yrTerm1.partialUpfront": "1.4335",
            "yrTerm3.partialUpfront": "1.0268"
          },
          "ondemand": "2.369"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.2466",
            "yrTerm3.allUpfront": "0.6193",
            "yrTerm1.allUpfront": "1.0442",
            "yrTerm1.partialUpfront": "1.0655",
            "yrTerm3.partialUpfront": "0.6588"
          },
          "ondemand": "2.001"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.664",
            "yrTerm3.allUpfront": "0.963",
            "yrTerm1.allUpfront": "1.3939",
            "yrTerm1.partialUpfront": "1.4225",
            "yrTerm3.partialUpfront": "1.0245"
          },
          "ondemand": "2.585"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.213",
            "yrTerm3.allUpfront": "1.403",
            "yrTerm1.allUpfront": "1.8532",
            "yrTerm1.partialUpfront": "1.8911",
            "yrTerm3.partialUpfront": "1.4925"
          },
          "ondemand": "4.254"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.628",
            "yrTerm3.allUpfront": "0.9337",
            "yrTerm1.allUpfront": "1.3636",
            "yrTerm1.partialUpfront": "1.3914",
            "yrTerm3.partialUpfront": "0.9933"
          },
          "ondemand": "2.337"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.22",
            "yrTerm3.allUpfront": "0.6084",
            "yrTerm1.allUpfront": "1.0255",
            "yrTerm1.partialUpfront": "1.0464",
            "yrTerm3.partialUpfront": "0.6473"
          },
          "ondemand": "2.035"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.664",
            "yrTerm3.allUpfront": "0.963",
            "yrTerm1.allUpfront": "1.3939",
            "yrTerm1.partialUpfront": "1.4225",
            "yrTerm3.partialUpfront": "1.0245"
          },
          "ondemand": "2.585"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "2.213",
            "yrTerm3.allUpfront": "1.403",
            "yrTerm1.allUpfront": "1.8532",
            "yrTerm1.partialUpfront": "1.8911",
            "yrTerm3.partialUpfront": "1.4925"
          },
          "ondemand": "4.254"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.628",
            "yrTerm3.allUpfront": "0.9337",
            "yrTerm1.allUpfront": "1.3636",
            "yrTerm1.partialUpfront": "1.3914",
            "yrTerm3.partialUpfront": "0.9933"
          },
          "ondemand": "2.337"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.22",
            "yrTerm3.allUpfront": "0.6084",
            "yrTerm1.allUpfront": "1.0255",
            "yrTerm1.partialUpfront": "1.0464",
            "yrTerm3.partialUpfront": "0.6473"
          },
          "ondemand": "2.035"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.155",
            "yrTerm3.allUpfront": "0.6688",
            "yrTerm1.allUpfront": "0.9678",
            "yrTerm1.partialUpfront": "0.988",
            "yrTerm3.partialUpfront": "0.7115"
          },
          "ondemand": "1.986"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.536",
            "yrTerm3.allUpfront": "0.974",
            "yrTerm1.allUpfront": "1.2868",
            "yrTerm1.partialUpfront": "1.313",
            "yrTerm3.partialUpfront": "1.0362"
          },
          "ondemand": "2.459"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.131",
            "yrTerm3.allUpfront": "0.6481",
            "yrTerm1.allUpfront": "0.9471",
            "yrTerm1.partialUpfront": "0.9665",
            "yrTerm3.partialUpfront": "0.6895"
          },
          "ondemand": "1.946"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.848",
            "yrTerm3.allUpfront": "0.4225",
            "yrTerm1.allUpfront": "0.712",
            "yrTerm1.partialUpfront": "0.7265",
            "yrTerm3.partialUpfront": "0.4495"
          },
          "ondemand": "1.705"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.326",
            "yrTerm3.allUpfront": "0.7534",
            "yrTerm1.allUpfront": "1.111",
            "yrTerm1.partialUpfront": "1.1337",
            "yrTerm3.partialUpfront": "0.8015"
          },
          "ondemand": "2.327"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.707",
            "yrTerm3.allUpfront": "1.0591",
            "yrTerm1.allUpfront": "1.4299",
            "yrTerm1.partialUpfront": "1.4591",
            "yrTerm3.partialUpfront": "1.1268"
          },
          "ondemand": "2.8"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.301",
            "yrTerm3.allUpfront": "0.733",
            "yrTerm1.allUpfront": "1.0896",
            "yrTerm1.partialUpfront": "1.1119",
            "yrTerm3.partialUpfront": "0.7798"
          },
          "ondemand": "2.287"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.018",
            "yrTerm3.allUpfront": "0.507",
            "yrTerm1.allUpfront": "0.8543",
            "yrTerm1.partialUpfront": "0.8716",
            "yrTerm3.partialUpfront": "0.5392"
          },
          "ondemand": "2.046"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.155",
            "yrTerm3.allUpfront": "0.6688",
            "yrTerm1.allUpfront": "0.9678",
            "yrTerm1.partialUpfront": "0.988",
            "yrTerm3.partialUpfront": "0.7115"
          },
          "ondemand": "2.156"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.92",
            "yrTerm3.allUpfront": "1.2175",
            "yrTerm1.allUpfront": "1.6082",
            "yrTerm1.partialUpfront": "1.641",
            "yrTerm3.partialUpfront": "1.2952"
          },
          "ondemand": "2.63"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.414",
            "yrTerm3.allUpfront": "0.8108",
            "yrTerm1.allUpfront": "1.1844",
            "yrTerm1.partialUpfront": "1.2085",
            "yrTerm3.partialUpfront": "0.8625"
          },
          "ondemand": "2.116"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.03",
            "yrTerm3.allUpfront": "0.5447",
            "yrTerm1.allUpfront": "0.8638",
            "yrTerm1.partialUpfront": "0.8815",
            "yrTerm3.partialUpfront": "0.5795"
          },
          "ondemand": "1.876"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.42",
            "yrTerm3.allUpfront": "0.8348",
            "yrTerm1.allUpfront": "1.1895",
            "yrTerm1.partialUpfront": "1.2138",
            "yrTerm3.partialUpfront": "0.8881"
          },
          "ondemand": "3.223"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.105",
            "yrTerm3.allUpfront": "2.0906",
            "yrTerm1.allUpfront": "2.601",
            "yrTerm1.partialUpfront": "2.6541",
            "yrTerm3.partialUpfront": "2.224"
          },
          "ondemand": "4.931"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.418",
            "yrTerm3.allUpfront": "0.8139",
            "yrTerm1.allUpfront": "1.1878",
            "yrTerm1.partialUpfront": "1.212",
            "yrTerm3.partialUpfront": "0.8659"
          },
          "ondemand": "2.267"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.112",
            "yrTerm3.allUpfront": "0.5883",
            "yrTerm1.allUpfront": "0.9329",
            "yrTerm1.partialUpfront": "0.952",
            "yrTerm3.partialUpfront": "0.6259"
          },
          "ondemand": "2.026"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "1.155",
            "yrTerm3.allUpfront": "0.6688",
            "yrTerm1.allUpfront": "0.9678",
            "yrTerm1.partialUpfront": "0.988",
            "yrTerm3.partialUpfront": "0.7115"
          },
          "ondemand": "1.986"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "1.92",
            "yrTerm3.allUpfront": "1.2175",
            "yrTerm1.allUpfront": "1.6082",
            "yrTerm1.partialUpfront": "1.641",
            "yrTerm3.partialUpfront": "1.2952"
          },
          "ondemand": "2.459"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.414",
            "yrTerm3.allUpfront": "0.8108",
            "yrTerm1.allUpfront": "1.1844",
            "yrTerm1.partialUpfront": "1.2085",
            "yrTerm3.partialUpfront": "0.8625"
          },
          "ondemand": "1.946"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.03",
            "yrTerm3.allUpfront": "0.5447",
            "yrTerm1.allUpfront": "0.8638",
            "yrTerm1.partialUpfront": "0.8815",
            "yrTerm3.partialUpfront": "0.5795"
          },
          "ondemand": "1.876"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 15,
      "max_enis": 4
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": true,
    "storage": {
      "ssd": true,
      "devices": 2,
      "size": 800.0
    },
    "instance_type": "i2.2xlarge",
    "ECU": 27.0,
    "memory": 61.0,
    "ebs_max_bandwidth": 1000.0
  },
  {
    "family": "Storage optimized",
    "enhanced_networking": true,
    "vCPU": 16,
    "generation": "current",
    "ebs_iops": 250.0,
    "network_performance": "High",
    "ebs_throughput": 16000.0,
    "pretty_name": "I2 Quadruple Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "2.311",
            "yrTerm3.allUpfront": "1.3366",
            "yrTerm1.allUpfront": "1.9355",
            "yrTerm1.partialUpfront": "1.9758",
            "yrTerm3.partialUpfront": "1.4219"
          },
          "ondemand": "3.971"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.071",
            "yrTerm3.allUpfront": "1.9472",
            "yrTerm1.allUpfront": "2.5724",
            "yrTerm1.partialUpfront": "2.6249",
            "yrTerm3.partialUpfront": "2.0715"
          },
          "ondemand": "4.918"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.26",
            "yrTerm3.allUpfront": "1.2963",
            "yrTerm1.allUpfront": "1.8934",
            "yrTerm1.partialUpfront": "1.9321",
            "yrTerm3.partialUpfront": "1.379"
          },
          "ondemand": "3.891"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.696",
            "yrTerm3.allUpfront": "0.8451",
            "yrTerm1.allUpfront": "1.4229",
            "yrTerm1.partialUpfront": "1.4521",
            "yrTerm3.partialUpfront": "0.899"
          },
          "ondemand": "3.41"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.39",
            "yrTerm3.allUpfront": "1.9605",
            "yrTerm1.allUpfront": "2.8398",
            "yrTerm1.partialUpfront": "2.898",
            "yrTerm3.partialUpfront": "2.0857"
          },
          "ondemand": "4.923"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "4.504",
            "yrTerm3.allUpfront": "2.8558",
            "yrTerm1.allUpfront": "3.7728",
            "yrTerm1.partialUpfront": "3.8498",
            "yrTerm3.partialUpfront": "3.0381"
          },
          "ondemand": "8.103"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "3.316",
            "yrTerm3.allUpfront": "1.9003",
            "yrTerm1.allUpfront": "2.7772",
            "yrTerm1.partialUpfront": "2.8338",
            "yrTerm3.partialUpfront": "2.0216"
          },
          "ondemand": "4.451"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.492",
            "yrTerm3.allUpfront": "1.2395",
            "yrTerm1.allUpfront": "2.0862",
            "yrTerm1.partialUpfront": "2.1288",
            "yrTerm3.partialUpfront": "1.3186"
          },
          "ondemand": "4.002"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.5187",
            "yrTerm3.allUpfront": "2.2675",
            "yrTerm1.allUpfront": "3.1143",
            "yrTerm1.partialUpfront": "3.1568",
            "yrTerm3.partialUpfront": "2.3466"
          },
          "ondemand": "5.03"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "6.0587",
            "yrTerm3.allUpfront": "4.8075",
            "yrTerm1.allUpfront": "5.6543",
            "yrTerm1.partialUpfront": "5.6968",
            "yrTerm3.partialUpfront": "4.8866"
          },
          "ondemand": "7.57"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "3.2267",
            "yrTerm3.allUpfront": "1.9755",
            "yrTerm1.allUpfront": "2.8223",
            "yrTerm1.partialUpfront": "2.8648",
            "yrTerm3.partialUpfront": "2.0546"
          },
          "ondemand": "4.738"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.4907",
            "yrTerm3.allUpfront": "1.2395",
            "yrTerm1.allUpfront": "2.0863",
            "yrTerm1.partialUpfront": "2.1288",
            "yrTerm3.partialUpfront": "1.3186"
          },
          "ondemand": "4.002"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.328",
            "yrTerm3.allUpfront": "1.925",
            "yrTerm1.allUpfront": "2.7878",
            "yrTerm1.partialUpfront": "2.8449",
            "yrTerm3.partialUpfront": "2.0479"
          },
          "ondemand": "5.169"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "4.422",
            "yrTerm3.allUpfront": "2.804",
            "yrTerm1.allUpfront": "3.7042",
            "yrTerm1.partialUpfront": "3.7799",
            "yrTerm3.partialUpfront": "2.983"
          },
          "ondemand": "8.508"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "3.255",
            "yrTerm3.allUpfront": "1.8664",
            "yrTerm1.allUpfront": "2.726",
            "yrTerm1.partialUpfront": "2.7817",
            "yrTerm3.partialUpfront": "1.9855"
          },
          "ondemand": "4.674"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.44",
            "yrTerm3.allUpfront": "1.2169",
            "yrTerm1.allUpfront": "2.0489",
            "yrTerm1.partialUpfront": "2.0907",
            "yrTerm3.partialUpfront": "1.2945"
          },
          "ondemand": "4.07"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "3.328",
            "yrTerm3.allUpfront": "1.925",
            "yrTerm1.allUpfront": "2.7878",
            "yrTerm1.partialUpfront": "2.8449",
            "yrTerm3.partialUpfront": "2.0479"
          },
          "ondemand": "5.169"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "4.422",
            "yrTerm3.allUpfront": "2.804",
            "yrTerm1.allUpfront": "3.7042",
            "yrTerm1.partialUpfront": "3.7799",
            "yrTerm3.partialUpfront": "2.983"
          },
          "ondemand": "8.508"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "3.255",
            "yrTerm3.allUpfront": "1.8664",
            "yrTerm1.allUpfront": "2.726",
            "yrTerm1.partialUpfront": "2.7817",
            "yrTerm3.partialUpfront": "1.9855"
          },
          "ondemand": "4.674"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.44",
            "yrTerm3.allUpfront": "1.2169",
            "yrTerm1.allUpfront": "2.0489",
            "yrTerm1.partialUpfront": "2.0907",
            "yrTerm3.partialUpfront": "1.2945"
          },
          "ondemand": "4.07"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "2.311",
            "yrTerm3.allUpfront": "1.3366",
            "yrTerm1.allUpfront": "1.9355",
            "yrTerm1.partialUpfront": "1.9758",
            "yrTerm3.partialUpfront": "1.4219"
          },
          "ondemand": "3.971"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.071",
            "yrTerm3.allUpfront": "1.9472",
            "yrTerm1.allUpfront": "2.5724",
            "yrTerm1.partialUpfront": "2.6249",
            "yrTerm3.partialUpfront": "2.0715"
          },
          "ondemand": "4.918"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.26",
            "yrTerm3.allUpfront": "1.2963",
            "yrTerm1.allUpfront": "1.8934",
            "yrTerm1.partialUpfront": "1.9321",
            "yrTerm3.partialUpfront": "1.379"
          },
          "ondemand": "3.891"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.696",
            "yrTerm3.allUpfront": "0.8451",
            "yrTerm1.allUpfront": "1.4229",
            "yrTerm1.partialUpfront": "1.4521",
            "yrTerm3.partialUpfront": "0.899"
          },
          "ondemand": "3.41"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "2.653",
            "yrTerm3.allUpfront": "1.5064",
            "yrTerm1.allUpfront": "2.2217",
            "yrTerm1.partialUpfront": "2.2671",
            "yrTerm3.partialUpfront": "1.6026"
          },
          "ondemand": "4.653"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.411",
            "yrTerm3.allUpfront": "2.1169",
            "yrTerm1.allUpfront": "2.8573",
            "yrTerm1.partialUpfront": "2.9156",
            "yrTerm3.partialUpfront": "2.2521"
          },
          "ondemand": "5.6"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.601",
            "yrTerm3.allUpfront": "1.466",
            "yrTerm1.allUpfront": "2.1787",
            "yrTerm1.partialUpfront": "2.2232",
            "yrTerm3.partialUpfront": "1.5596"
          },
          "ondemand": "4.573"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.036",
            "yrTerm3.allUpfront": "1.0141",
            "yrTerm1.allUpfront": "1.7075",
            "yrTerm1.partialUpfront": "1.7423",
            "yrTerm3.partialUpfront": "1.0785"
          },
          "ondemand": "4.092"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "2.311",
            "yrTerm3.allUpfront": "1.3366",
            "yrTerm1.allUpfront": "1.9355",
            "yrTerm1.partialUpfront": "1.9758",
            "yrTerm3.partialUpfront": "1.4219"
          },
          "ondemand": "4.312"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.839",
            "yrTerm3.allUpfront": "2.4341",
            "yrTerm1.allUpfront": "3.2153",
            "yrTerm1.partialUpfront": "3.2809",
            "yrTerm3.partialUpfront": "2.5895"
          },
          "ondemand": "5.259"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.827",
            "yrTerm3.allUpfront": "1.6206",
            "yrTerm1.allUpfront": "2.3677",
            "yrTerm1.partialUpfront": "2.4161",
            "yrTerm3.partialUpfront": "1.724"
          },
          "ondemand": "4.232"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.06",
            "yrTerm3.allUpfront": "1.0895",
            "yrTerm1.allUpfront": "1.7267",
            "yrTerm1.partialUpfront": "1.7621",
            "yrTerm3.partialUpfront": "1.159"
          },
          "ondemand": "3.751"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "2.839",
            "yrTerm3.allUpfront": "1.6678",
            "yrTerm1.allUpfront": "2.378",
            "yrTerm1.partialUpfront": "2.4265",
            "yrTerm3.partialUpfront": "1.7743"
          },
          "ondemand": "5.402"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "5.425",
            "yrTerm3.allUpfront": "3.6174",
            "yrTerm1.allUpfront": "4.5436",
            "yrTerm1.partialUpfront": "4.6363",
            "yrTerm3.partialUpfront": "3.8482"
          },
          "ondemand": "8.608"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.835",
            "yrTerm3.allUpfront": "1.6278",
            "yrTerm1.allUpfront": "2.3748",
            "yrTerm1.partialUpfront": "2.4232",
            "yrTerm3.partialUpfront": "1.7317"
          },
          "ondemand": "4.532"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.224",
            "yrTerm3.allUpfront": "1.1766",
            "yrTerm1.allUpfront": "1.8648",
            "yrTerm1.partialUpfront": "1.9029",
            "yrTerm3.partialUpfront": "1.2517"
          },
          "ondemand": "4.051"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "2.311",
            "yrTerm3.allUpfront": "1.3366",
            "yrTerm1.allUpfront": "1.9355",
            "yrTerm1.partialUpfront": "1.9758",
            "yrTerm3.partialUpfront": "1.4219"
          },
          "ondemand": "3.971"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "3.839",
            "yrTerm3.allUpfront": "2.4341",
            "yrTerm1.allUpfront": "3.2153",
            "yrTerm1.partialUpfront": "3.2809",
            "yrTerm3.partialUpfront": "2.5895"
          },
          "ondemand": "4.918"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.827",
            "yrTerm3.allUpfront": "1.6206",
            "yrTerm1.allUpfront": "2.3677",
            "yrTerm1.partialUpfront": "2.4161",
            "yrTerm3.partialUpfront": "1.724"
          },
          "ondemand": "3.891"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.06",
            "yrTerm3.allUpfront": "1.0895",
            "yrTerm1.allUpfront": "1.7267",
            "yrTerm1.partialUpfront": "1.7621",
            "yrTerm3.partialUpfront": "1.159"
          },
          "ondemand": "3.751"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 30,
      "max_enis": 8
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": true,
    "storage": {
      "ssd": true,
      "devices": 4,
      "size": 800.0
    },
    "instance_type": "i2.4xlarge",
    "ECU": 53.0,
    "memory": 122.0,
    "ebs_max_bandwidth": 2000.0
  },
  {
    "family": "Storage optimized",
    "enhanced_networking": true,
    "vCPU": 32,
    "generation": "current",
    "ebs_iops": 0,
    "network_performance": "10 Gigabit",
    "ebs_throughput": 0,
    "pretty_name": "I2 Eight Extra Large",
    "pricing": {
      "us-east-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "4.623",
            "yrTerm3.allUpfront": "2.6732",
            "yrTerm1.allUpfront": "3.872",
            "yrTerm1.partialUpfront": "3.9516",
            "yrTerm3.partialUpfront": "2.8438"
          },
          "ondemand": "7.942"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "6.142",
            "yrTerm3.allUpfront": "3.8944",
            "yrTerm1.allUpfront": "5.1447",
            "yrTerm1.partialUpfront": "5.2498",
            "yrTerm3.partialUpfront": "4.1429"
          },
          "ondemand": "9.836"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "4.521",
            "yrTerm3.allUpfront": "2.5926",
            "yrTerm1.allUpfront": "3.7869",
            "yrTerm1.partialUpfront": "3.8641",
            "yrTerm3.partialUpfront": "2.7581"
          },
          "ondemand": "7.782"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "3.392",
            "yrTerm3.allUpfront": "1.6902",
            "yrTerm1.allUpfront": "2.846",
            "yrTerm1.partialUpfront": "2.9041",
            "yrTerm3.partialUpfront": "1.7981"
          },
          "ondemand": "6.82"
        }
      },
      "ap-northeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "6.781",
            "yrTerm3.allUpfront": "3.9201",
            "yrTerm1.allUpfront": "5.6798",
            "yrTerm1.partialUpfront": "5.7961",
            "yrTerm3.partialUpfront": "4.1703"
          },
          "ondemand": "9.846"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "9.009",
            "yrTerm3.allUpfront": "5.7116",
            "yrTerm1.allUpfront": "7.5457",
            "yrTerm1.partialUpfront": "7.6997",
            "yrTerm3.partialUpfront": "6.0762"
          },
          "ondemand": "16.206"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "6.631",
            "yrTerm3.allUpfront": "3.8026",
            "yrTerm1.allUpfront": "5.5543",
            "yrTerm1.partialUpfront": "5.6677",
            "yrTerm3.partialUpfront": "4.0454"
          },
          "ondemand": "8.903"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "4.984",
            "yrTerm3.allUpfront": "2.4791",
            "yrTerm1.allUpfront": "4.1744",
            "yrTerm1.partialUpfront": "4.2597",
            "yrTerm3.partialUpfront": "2.6374"
          },
          "ondemand": "8.004"
        }
      },
      "ap-northeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "7.0398",
            "yrTerm3.allUpfront": "4.5351",
            "yrTerm1.allUpfront": "6.2305",
            "yrTerm1.partialUpfront": "6.3157",
            "yrTerm3.partialUpfront": "4.6934"
          },
          "ondemand": "10.06"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "12.1198",
            "yrTerm3.allUpfront": "9.6151",
            "yrTerm1.allUpfront": "11.3105",
            "yrTerm1.partialUpfront": "11.3957",
            "yrTerm3.partialUpfront": "9.7734"
          },
          "ondemand": "15.14"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "6.4558",
            "yrTerm3.allUpfront": "3.9511",
            "yrTerm1.allUpfront": "5.6465",
            "yrTerm1.partialUpfront": "5.7317",
            "yrTerm3.partialUpfront": "4.1094"
          },
          "ondemand": "9.476"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "4.9838",
            "yrTerm3.allUpfront": "2.4791",
            "yrTerm1.allUpfront": "4.1745",
            "yrTerm1.partialUpfront": "4.2597",
            "yrTerm3.partialUpfront": "2.6374"
          },
          "ondemand": "8.004"
        }
      },
      "ap-southeast-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "6.657",
            "yrTerm3.allUpfront": "3.8491",
            "yrTerm1.allUpfront": "5.5756",
            "yrTerm1.partialUpfront": "5.6898",
            "yrTerm3.partialUpfront": "4.0948"
          },
          "ondemand": "10.338"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "8.845",
            "yrTerm3.allUpfront": "5.608",
            "yrTerm1.allUpfront": "7.4084",
            "yrTerm1.partialUpfront": "7.5596",
            "yrTerm3.partialUpfront": "5.966"
          },
          "ondemand": "17.016"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "6.51",
            "yrTerm3.allUpfront": "3.7338",
            "yrTerm1.allUpfront": "5.4531",
            "yrTerm1.partialUpfront": "5.5644",
            "yrTerm3.partialUpfront": "3.9721"
          },
          "ondemand": "9.348"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "4.88",
            "yrTerm3.allUpfront": "2.4338",
            "yrTerm1.allUpfront": "4.0977",
            "yrTerm1.partialUpfront": "4.1814",
            "yrTerm3.partialUpfront": "2.5891"
          },
          "ondemand": "8.14"
        }
      },
      "ap-southeast-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "6.657",
            "yrTerm3.allUpfront": "3.8491",
            "yrTerm1.allUpfront": "5.5756",
            "yrTerm1.partialUpfront": "5.6898",
            "yrTerm3.partialUpfront": "4.0948"
          },
          "ondemand": "10.338"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "8.845",
            "yrTerm3.allUpfront": "5.608",
            "yrTerm1.allUpfront": "7.4084",
            "yrTerm1.partialUpfront": "7.5596",
            "yrTerm3.partialUpfront": "5.966"
          },
          "ondemand": "17.016"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "6.51",
            "yrTerm3.allUpfront": "3.7338",
            "yrTerm1.allUpfront": "5.4531",
            "yrTerm1.partialUpfront": "5.5644",
            "yrTerm3.partialUpfront": "3.9721"
          },
          "ondemand": "9.348"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "4.88",
            "yrTerm3.allUpfront": "2.4338",
            "yrTerm1.allUpfront": "4.0977",
            "yrTerm1.partialUpfront": "4.1814",
            "yrTerm3.partialUpfront": "2.5891"
          },
          "ondemand": "8.14"
        }
      },
      "us-west-2": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "4.623",
            "yrTerm3.allUpfront": "2.6732",
            "yrTerm1.allUpfront": "3.872",
            "yrTerm1.partialUpfront": "3.9516",
            "yrTerm3.partialUpfront": "2.8438"
          },
          "ondemand": "7.942"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "6.142",
            "yrTerm3.allUpfront": "3.8944",
            "yrTerm1.allUpfront": "5.1447",
            "yrTerm1.partialUpfront": "5.2498",
            "yrTerm3.partialUpfront": "4.1429"
          },
          "ondemand": "9.836"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "4.521",
            "yrTerm3.allUpfront": "2.5926",
            "yrTerm1.allUpfront": "3.7869",
            "yrTerm1.partialUpfront": "3.8641",
            "yrTerm3.partialUpfront": "2.7581"
          },
          "ondemand": "7.782"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "3.392",
            "yrTerm3.allUpfront": "1.6902",
            "yrTerm1.allUpfront": "2.846",
            "yrTerm1.partialUpfront": "2.9041",
            "yrTerm3.partialUpfront": "1.7981"
          },
          "ondemand": "6.82"
        }
      },
      "us-gov-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "5.304",
            "yrTerm3.allUpfront": "3.0119",
            "yrTerm1.allUpfront": "4.4425",
            "yrTerm1.partialUpfront": "4.5332",
            "yrTerm3.partialUpfront": "3.2042"
          },
          "ondemand": "9.306"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "6.823",
            "yrTerm3.allUpfront": "4.233",
            "yrTerm1.allUpfront": "5.7146",
            "yrTerm1.partialUpfront": "5.8313",
            "yrTerm3.partialUpfront": "4.5031"
          },
          "ondemand": "11.2"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "5.201",
            "yrTerm3.allUpfront": "2.9311",
            "yrTerm1.allUpfront": "4.3565",
            "yrTerm1.partialUpfront": "4.4455",
            "yrTerm3.partialUpfront": "3.1182"
          },
          "ondemand": "9.146"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "4.072",
            "yrTerm3.allUpfront": "2.0282",
            "yrTerm1.allUpfront": "3.4152",
            "yrTerm1.partialUpfront": "3.4845",
            "yrTerm3.partialUpfront": "2.1569"
          },
          "ondemand": "8.184"
        }
      },
      "us-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "4.623",
            "yrTerm3.allUpfront": "2.6732",
            "yrTerm1.allUpfront": "3.872",
            "yrTerm1.partialUpfront": "3.9516",
            "yrTerm3.partialUpfront": "2.8438"
          },
          "ondemand": "8.624"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "7.677",
            "yrTerm3.allUpfront": "4.8682",
            "yrTerm1.allUpfront": "6.4305",
            "yrTerm1.partialUpfront": "6.5618",
            "yrTerm3.partialUpfront": "5.1789"
          },
          "ondemand": "10.518"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "5.654",
            "yrTerm3.allUpfront": "3.2412",
            "yrTerm1.allUpfront": "4.7355",
            "yrTerm1.partialUpfront": "4.8321",
            "yrTerm3.partialUpfront": "3.4481"
          },
          "ondemand": "8.464"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "4.12",
            "yrTerm3.allUpfront": "2.179",
            "yrTerm1.allUpfront": "3.4535",
            "yrTerm1.partialUpfront": "3.5241",
            "yrTerm3.partialUpfront": "2.3181"
          },
          "ondemand": "7.502"
        }
      },
      "eu-central-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "5.678",
            "yrTerm3.allUpfront": "3.3366",
            "yrTerm1.allUpfront": "4.7559",
            "yrTerm1.partialUpfront": "4.853",
            "yrTerm3.partialUpfront": "3.5495"
          },
          "ondemand": "9.775"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "10.362",
            "yrTerm3.allUpfront": "6.8849",
            "yrTerm1.allUpfront": "8.679",
            "yrTerm1.partialUpfront": "8.8561",
            "yrTerm3.partialUpfront": "7.3244"
          },
          "ondemand": "16.439"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "5.67",
            "yrTerm3.allUpfront": "3.2556",
            "yrTerm1.allUpfront": "4.7495",
            "yrTerm1.partialUpfront": "4.8465",
            "yrTerm3.partialUpfront": "3.4634"
          },
          "ondemand": "9.064"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "4.448",
            "yrTerm3.allUpfront": "2.3533",
            "yrTerm1.allUpfront": "3.7298",
            "yrTerm1.partialUpfront": "3.8059",
            "yrTerm3.partialUpfront": "2.5035"
          },
          "ondemand": "8.102"
        }
      },
      "eu-west-1": {
        "mswinSQLWeb": {
          "reserved": {
            "yrTerm1.noUpfront": "4.623",
            "yrTerm3.allUpfront": "2.6732",
            "yrTerm1.allUpfront": "3.872",
            "yrTerm1.partialUpfront": "3.9516",
            "yrTerm3.partialUpfront": "2.8438"
          },
          "ondemand": "7.942"
        },
        "mswinSQL": {
          "reserved": {
            "yrTerm1.noUpfront": "7.677",
            "yrTerm3.allUpfront": "4.8682",
            "yrTerm1.allUpfront": "6.4305",
            "yrTerm1.partialUpfront": "6.5618",
            "yrTerm3.partialUpfront": "5.1789"
          },
          "ondemand": "9.836"
        },
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "5.654",
            "yrTerm3.allUpfront": "3.2412",
            "yrTerm1.allUpfront": "4.7355",
            "yrTerm1.partialUpfront": "4.8321",
            "yrTerm3.partialUpfront": "3.4481"
          },
          "ondemand": "7.782"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "4.12",
            "yrTerm3.allUpfront": "2.179",
            "yrTerm1.allUpfront": "3.4535",
            "yrTerm1.partialUpfront": "3.5241",
            "yrTerm3.partialUpfront": "2.3181"
          },
          "ondemand": "7.502"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 30,
      "max_enis": 8
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": false,
    "storage": {
      "ssd": true,
      "devices": 8,
      "size": 800.0
    },
    "instance_type": "i2.8xlarge",
    "ECU": 104.0,
    "memory": 244.0,
    "ebs_max_bandwidth": 0
  },
  {
    "family": "Storage optimized",
    "enhanced_networking": true,
    "vCPU": 4,
    "generation": "current",
    "ebs_iops": 93.75,
    "network_performance": "Moderate",
    "ebs_throughput": 6000.0,
    "pretty_name": "D2 Extra Large",
    "pricing": {
      "us-east-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.472",
            "yrTerm3.allUpfront": "0.2544",
            "yrTerm1.allUpfront": "0.3957",
            "yrTerm1.partialUpfront": "0.4038",
            "yrTerm3.partialUpfront": "0.2703"
          },
          "ondemand": "0.821"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.402",
            "yrTerm3.allUpfront": "0.1977",
            "yrTerm1.allUpfront": "0.337",
            "yrTerm1.partialUpfront": "0.3439",
            "yrTerm3.partialUpfront": "0.2103"
          },
          "ondemand": "0.69"
        }
      },
      "ap-northeast-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.599",
            "yrTerm3.allUpfront": "0.359",
            "yrTerm1.allUpfront": "0.5016",
            "yrTerm1.partialUpfront": "0.5119",
            "yrTerm3.partialUpfront": "0.3819"
          },
          "ondemand": "0.975"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.529",
            "yrTerm3.allUpfront": "0.3026",
            "yrTerm1.allUpfront": "0.4429",
            "yrTerm1.partialUpfront": "0.4519",
            "yrTerm3.partialUpfront": "0.322"
          },
          "ondemand": "0.844"
        }
      },
      "ap-northeast-2": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.7127",
            "yrTerm3.allUpfront": "0.4866",
            "yrTerm1.allUpfront": "0.6269",
            "yrTerm1.partialUpfront": "0.6359",
            "yrTerm3.partialUpfront": "0.506"
          },
          "ondemand": "1.028"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.5287",
            "yrTerm3.allUpfront": "0.3026",
            "yrTerm1.allUpfront": "0.4429",
            "yrTerm1.partialUpfront": "0.4519",
            "yrTerm3.partialUpfront": "0.322"
          },
          "ondemand": "0.844"
        }
      },
      "ap-southeast-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.584",
            "yrTerm3.allUpfront": "0.3502",
            "yrTerm1.allUpfront": "0.4888",
            "yrTerm1.partialUpfront": "0.4984",
            "yrTerm3.partialUpfront": "0.3723"
          },
          "ondemand": "1.001"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.513",
            "yrTerm3.allUpfront": "0.2935",
            "yrTerm1.allUpfront": "0.4297",
            "yrTerm1.partialUpfront": "0.4384",
            "yrTerm3.partialUpfront": "0.3122"
          },
          "ondemand": "0.87"
        }
      },
      "ap-southeast-2": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.584",
            "yrTerm3.allUpfront": "0.3502",
            "yrTerm1.allUpfront": "0.4888",
            "yrTerm1.partialUpfront": "0.4984",
            "yrTerm3.partialUpfront": "0.3723"
          },
          "ondemand": "1.001"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.513",
            "yrTerm3.allUpfront": "0.2935",
            "yrTerm1.allUpfront": "0.4297",
            "yrTerm1.partialUpfront": "0.4384",
            "yrTerm3.partialUpfront": "0.3122"
          },
          "ondemand": "0.87"
        }
      },
      "us-west-2": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.472",
            "yrTerm3.allUpfront": "0.2544",
            "yrTerm1.allUpfront": "0.3957",
            "yrTerm1.partialUpfront": "0.4038",
            "yrTerm3.partialUpfront": "0.2703"
          },
          "ondemand": "0.821"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.402",
            "yrTerm3.allUpfront": "0.1977",
            "yrTerm1.allUpfront": "0.337",
            "yrTerm1.partialUpfront": "0.3439",
            "yrTerm3.partialUpfront": "0.2103"
          },
          "ondemand": "0.69"
        }
      },
      "us-gov-west-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.633",
            "yrTerm3.allUpfront": "0.387",
            "yrTerm1.allUpfront": "0.555",
            "yrTerm1.partialUpfront": "0.564",
            "yrTerm3.partialUpfront": "0.402"
          },
          "ondemand": "0.976"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.483",
            "yrTerm3.allUpfront": "0.237",
            "yrTerm1.allUpfront": "0.405",
            "yrTerm1.partialUpfront": "0.414",
            "yrTerm3.partialUpfront": "0.252"
          },
          "ondemand": "0.828"
        }
      },
      "eu-central-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.553",
            "yrTerm3.allUpfront": "0.332",
            "yrTerm1.allUpfront": "0.4628",
            "yrTerm1.partialUpfront": "0.4721",
            "yrTerm3.partialUpfront": "0.3536"
          },
          "ondemand": "0.925"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.482",
            "yrTerm3.allUpfront": "0.276",
            "yrTerm1.allUpfront": "0.404",
            "yrTerm1.partialUpfront": "0.4122",
            "yrTerm3.partialUpfront": "0.2936"
          },
          "ondemand": "0.794"
        }
      },
      "eu-west-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.517",
            "yrTerm3.allUpfront": "0.3116",
            "yrTerm1.allUpfront": "0.4329",
            "yrTerm1.partialUpfront": "0.4419",
            "yrTerm3.partialUpfront": "0.3318"
          },
          "ondemand": "0.866"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.447",
            "yrTerm3.allUpfront": "0.2554",
            "yrTerm1.allUpfront": "0.3742",
            "yrTerm1.partialUpfront": "0.3819",
            "yrTerm3.partialUpfront": "0.2717"
          },
          "ondemand": "0.735"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 15,
      "max_enis": 4
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": true,
    "storage": {
      "ssd": false,
      "devices": 3,
      "size": 2000.0
    },
    "instance_type": "d2.xlarge",
    "ECU": 14.0,
    "memory": 30.5,
    "ebs_max_bandwidth": 750.0
  },
  {
    "family": "Storage optimized",
    "enhanced_networking": true,
    "vCPU": 8,
    "generation": "current",
    "ebs_iops": 125.0,
    "network_performance": "High",
    "ebs_throughput": 8000.0,
    "pretty_name": "D2 Double Extra Large",
    "pricing": {
      "us-east-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.885",
            "yrTerm3.allUpfront": "0.4602",
            "yrTerm1.allUpfront": "0.7417",
            "yrTerm1.partialUpfront": "0.7564",
            "yrTerm3.partialUpfront": "0.4898"
          },
          "ondemand": "1.601"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.804",
            "yrTerm3.allUpfront": "0.3954",
            "yrTerm1.allUpfront": "0.674",
            "yrTerm1.partialUpfront": "0.6878",
            "yrTerm3.partialUpfront": "0.4206"
          },
          "ondemand": "1.38"
        }
      },
      "ap-northeast-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.138",
            "yrTerm3.allUpfront": "0.6701",
            "yrTerm1.allUpfront": "0.9534",
            "yrTerm1.partialUpfront": "0.9724",
            "yrTerm3.partialUpfront": "0.7125"
          },
          "ondemand": "1.909"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.058",
            "yrTerm3.allUpfront": "0.6053",
            "yrTerm1.allUpfront": "0.8858",
            "yrTerm1.partialUpfront": "0.9038",
            "yrTerm3.partialUpfront": "0.6439"
          },
          "ondemand": "1.688"
        }
      },
      "ap-northeast-2": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.4255",
            "yrTerm3.allUpfront": "0.9733",
            "yrTerm1.allUpfront": "1.2537",
            "yrTerm1.partialUpfront": "1.2718",
            "yrTerm3.partialUpfront": "1.0119"
          },
          "ondemand": "2.056"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.0575",
            "yrTerm3.allUpfront": "0.6053",
            "yrTerm1.allUpfront": "0.8857",
            "yrTerm1.partialUpfront": "0.9038",
            "yrTerm3.partialUpfront": "0.6439"
          },
          "ondemand": "1.688"
        }
      },
      "ap-southeast-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.107",
            "yrTerm3.allUpfront": "0.6519",
            "yrTerm1.allUpfront": "0.9269",
            "yrTerm1.partialUpfront": "0.9459",
            "yrTerm3.partialUpfront": "0.6937"
          },
          "ondemand": "1.961"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.026",
            "yrTerm3.allUpfront": "0.587",
            "yrTerm1.allUpfront": "0.8594",
            "yrTerm1.partialUpfront": "0.8768",
            "yrTerm3.partialUpfront": "0.6245"
          },
          "ondemand": "1.74"
        }
      },
      "ap-southeast-2": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.107",
            "yrTerm3.allUpfront": "0.6519",
            "yrTerm1.allUpfront": "0.9269",
            "yrTerm1.partialUpfront": "0.9459",
            "yrTerm3.partialUpfront": "0.6937"
          },
          "ondemand": "1.961"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.026",
            "yrTerm3.allUpfront": "0.587",
            "yrTerm1.allUpfront": "0.8594",
            "yrTerm1.partialUpfront": "0.8768",
            "yrTerm3.partialUpfront": "0.6245"
          },
          "ondemand": "1.74"
        }
      },
      "us-west-2": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.885",
            "yrTerm3.allUpfront": "0.4602",
            "yrTerm1.allUpfront": "0.7417",
            "yrTerm1.partialUpfront": "0.7564",
            "yrTerm3.partialUpfront": "0.4898"
          },
          "ondemand": "1.601"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.804",
            "yrTerm3.allUpfront": "0.3954",
            "yrTerm1.allUpfront": "0.674",
            "yrTerm1.partialUpfront": "0.6878",
            "yrTerm3.partialUpfront": "0.4206"
          },
          "ondemand": "1.38"
        }
      },
      "us-gov-west-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.265",
            "yrTerm3.allUpfront": "0.775",
            "yrTerm1.allUpfront": "1.108",
            "yrTerm1.partialUpfront": "1.126",
            "yrTerm3.partialUpfront": "0.806"
          },
          "ondemand": "1.952"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.965",
            "yrTerm3.allUpfront": "0.475",
            "yrTerm1.allUpfront": "0.808",
            "yrTerm1.partialUpfront": "0.826",
            "yrTerm3.partialUpfront": "0.506"
          },
          "ondemand": "1.656"
        }
      },
      "eu-central-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.045",
            "yrTerm3.allUpfront": "0.6169",
            "yrTerm1.allUpfront": "0.8755",
            "yrTerm1.partialUpfront": "0.8937",
            "yrTerm3.partialUpfront": "0.6561"
          },
          "ondemand": "1.809"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.964",
            "yrTerm3.allUpfront": "0.552",
            "yrTerm1.allUpfront": "0.808",
            "yrTerm1.partialUpfront": "0.8243",
            "yrTerm3.partialUpfront": "0.5872"
          },
          "ondemand": "1.588"
        }
      },
      "eu-west-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "0.974",
            "yrTerm3.allUpfront": "0.5757",
            "yrTerm1.allUpfront": "0.8161",
            "yrTerm1.partialUpfront": "0.8323",
            "yrTerm3.partialUpfront": "0.6122"
          },
          "ondemand": "1.691"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "0.894",
            "yrTerm3.allUpfront": "0.5109",
            "yrTerm1.allUpfront": "0.7484",
            "yrTerm1.partialUpfront": "0.7637",
            "yrTerm3.partialUpfront": "0.5435"
          },
          "ondemand": "1.47"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 15,
      "max_enis": 4
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": true,
    "storage": {
      "ssd": false,
      "devices": 6,
      "size": 2000.0
    },
    "instance_type": "d2.2xlarge",
    "ECU": 28.0,
    "memory": 61.0,
    "ebs_max_bandwidth": 1000.0
  },
  {
    "family": "Storage optimized",
    "enhanced_networking": true,
    "vCPU": 16,
    "generation": "current",
    "ebs_iops": 250.0,
    "network_performance": "High",
    "ebs_throughput": 16000.0,
    "pretty_name": "D2 Quadruple Extra Large",
    "pricing": {
      "us-east-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.69",
            "yrTerm3.allUpfront": "0.8556",
            "yrTerm1.allUpfront": "1.4158",
            "yrTerm1.partialUpfront": "1.4444",
            "yrTerm3.partialUpfront": "0.9101"
          },
          "ondemand": "3.062"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.608",
            "yrTerm3.allUpfront": "0.7907",
            "yrTerm1.allUpfront": "1.3479",
            "yrTerm1.partialUpfront": "1.3757",
            "yrTerm3.partialUpfront": "0.8412"
          },
          "ondemand": "2.76"
        }
      },
      "ap-northeast-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.196",
            "yrTerm3.allUpfront": "1.2754",
            "yrTerm1.allUpfront": "1.8392",
            "yrTerm1.partialUpfront": "1.8764",
            "yrTerm3.partialUpfront": "1.3564"
          },
          "ondemand": "3.678"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.116",
            "yrTerm3.allUpfront": "1.2105",
            "yrTerm1.allUpfront": "1.7717",
            "yrTerm1.partialUpfront": "1.8077",
            "yrTerm3.partialUpfront": "1.2878"
          },
          "ondemand": "3.376"
        }
      },
      "ap-northeast-2": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.851",
            "yrTerm3.allUpfront": "1.9466",
            "yrTerm1.allUpfront": "2.5075",
            "yrTerm1.partialUpfront": "2.5437",
            "yrTerm3.partialUpfront": "2.0238"
          },
          "ondemand": "4.112"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.115",
            "yrTerm3.allUpfront": "1.2106",
            "yrTerm1.allUpfront": "1.7715",
            "yrTerm1.partialUpfront": "1.8077",
            "yrTerm3.partialUpfront": "1.2878"
          },
          "ondemand": "3.376"
        }
      },
      "ap-southeast-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.132",
            "yrTerm3.allUpfront": "1.2389",
            "yrTerm1.allUpfront": "1.7862",
            "yrTerm1.partialUpfront": "1.8223",
            "yrTerm3.partialUpfront": "1.318"
          },
          "ondemand": "3.782"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.052",
            "yrTerm3.allUpfront": "1.174",
            "yrTerm1.allUpfront": "1.7187",
            "yrTerm1.partialUpfront": "1.7536",
            "yrTerm3.partialUpfront": "1.249"
          },
          "ondemand": "3.48"
        }
      },
      "ap-southeast-2": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.132",
            "yrTerm3.allUpfront": "1.2389",
            "yrTerm1.allUpfront": "1.7862",
            "yrTerm1.partialUpfront": "1.8223",
            "yrTerm3.partialUpfront": "1.318"
          },
          "ondemand": "3.782"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "2.052",
            "yrTerm3.allUpfront": "1.174",
            "yrTerm1.allUpfront": "1.7187",
            "yrTerm1.partialUpfront": "1.7536",
            "yrTerm3.partialUpfront": "1.249"
          },
          "ondemand": "3.48"
        }
      },
      "us-west-2": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.69",
            "yrTerm3.allUpfront": "0.8556",
            "yrTerm1.allUpfront": "1.4158",
            "yrTerm1.partialUpfront": "1.4444",
            "yrTerm3.partialUpfront": "0.9101"
          },
          "ondemand": "3.062"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.608",
            "yrTerm3.allUpfront": "0.7907",
            "yrTerm1.allUpfront": "1.3479",
            "yrTerm1.partialUpfront": "1.3757",
            "yrTerm3.partialUpfront": "0.8412"
          },
          "ondemand": "2.76"
        }
      },
      "us-gov-west-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.532",
            "yrTerm3.allUpfront": "1.548",
            "yrTerm1.allUpfront": "2.218",
            "yrTerm1.partialUpfront": "2.251",
            "yrTerm3.partialUpfront": "1.609"
          },
          "ondemand": "3.904"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.932",
            "yrTerm3.allUpfront": "0.948",
            "yrTerm1.allUpfront": "1.618",
            "yrTerm1.partialUpfront": "1.651",
            "yrTerm3.partialUpfront": "1.009"
          },
          "ondemand": "3.312"
        }
      },
      "eu-central-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "2.01",
            "yrTerm3.allUpfront": "1.1688",
            "yrTerm1.allUpfront": "1.6833",
            "yrTerm1.partialUpfront": "1.7178",
            "yrTerm3.partialUpfront": "1.2437"
          },
          "ondemand": "3.478"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.928",
            "yrTerm3.allUpfront": "1.104",
            "yrTerm1.allUpfront": "1.616",
            "yrTerm1.partialUpfront": "1.6487",
            "yrTerm3.partialUpfront": "1.1745"
          },
          "ondemand": "3.176"
        }
      },
      "eu-west-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "1.868",
            "yrTerm3.allUpfront": "1.0866",
            "yrTerm1.allUpfront": "1.5645",
            "yrTerm1.partialUpfront": "1.5963",
            "yrTerm3.partialUpfront": "1.156"
          },
          "ondemand": "3.242"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "1.788",
            "yrTerm3.allUpfront": "1.0218",
            "yrTerm1.allUpfront": "1.4968",
            "yrTerm1.partialUpfront": "1.5275",
            "yrTerm3.partialUpfront": "1.0869"
          },
          "ondemand": "2.94"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 30,
      "max_enis": 8
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": true,
    "storage": {
      "ssd": false,
      "devices": 12,
      "size": 2000.0
    },
    "instance_type": "d2.4xlarge",
    "ECU": 56.0,
    "memory": 122.0,
    "ebs_max_bandwidth": 2000.0
  },
  {
    "family": "Storage optimized",
    "enhanced_networking": true,
    "vCPU": 36,
    "generation": "current",
    "ebs_iops": 500.0,
    "network_performance": "10 Gigabit",
    "ebs_throughput": 32000.0,
    "pretty_name": "D2 Eight Extra Large",
    "pricing": {
      "us-east-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "3.3",
            "yrTerm3.allUpfront": "1.6462",
            "yrTerm1.allUpfront": "2.7639",
            "yrTerm1.partialUpfront": "2.8202",
            "yrTerm3.partialUpfront": "1.7516"
          },
          "ondemand": "6.198"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "3.216",
            "yrTerm3.allUpfront": "1.5814",
            "yrTerm1.allUpfront": "2.6959",
            "yrTerm1.partialUpfront": "2.7513",
            "yrTerm3.partialUpfront": "1.6823"
          },
          "ondemand": "5.52"
        }
      },
      "ap-northeast-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "4.311",
            "yrTerm3.allUpfront": "2.486",
            "yrTerm1.allUpfront": "3.6106",
            "yrTerm1.partialUpfront": "3.6841",
            "yrTerm3.partialUpfront": "2.6443"
          },
          "ondemand": "7.43"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "4.232",
            "yrTerm3.allUpfront": "2.421",
            "yrTerm1.allUpfront": "3.5434",
            "yrTerm1.partialUpfront": "3.6153",
            "yrTerm3.partialUpfront": "2.5757"
          },
          "ondemand": "6.752"
        }
      },
      "ap-northeast-2": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "5.8859",
            "yrTerm3.allUpfront": "4.0771",
            "yrTerm1.allUpfront": "5.199",
            "yrTerm1.partialUpfront": "5.2713",
            "yrTerm3.partialUpfront": "4.2317"
          },
          "ondemand": "8.408"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "4.2299",
            "yrTerm3.allUpfront": "2.4211",
            "yrTerm1.allUpfront": "3.543",
            "yrTerm1.partialUpfront": "3.6153",
            "yrTerm3.partialUpfront": "2.5757"
          },
          "ondemand": "6.752"
        }
      },
      "ap-southeast-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "4.184",
            "yrTerm3.allUpfront": "2.4129",
            "yrTerm1.allUpfront": "3.5047",
            "yrTerm1.partialUpfront": "3.5761",
            "yrTerm3.partialUpfront": "2.5664"
          },
          "ondemand": "7.638"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "4.104",
            "yrTerm3.allUpfront": "2.3479",
            "yrTerm1.allUpfront": "3.4374",
            "yrTerm1.partialUpfront": "3.5073",
            "yrTerm3.partialUpfront": "2.4979"
          },
          "ondemand": "6.96"
        }
      },
      "ap-southeast-2": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "4.184",
            "yrTerm3.allUpfront": "2.4129",
            "yrTerm1.allUpfront": "3.5047",
            "yrTerm1.partialUpfront": "3.5761",
            "yrTerm3.partialUpfront": "2.5664"
          },
          "ondemand": "7.638"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "4.104",
            "yrTerm3.allUpfront": "2.3479",
            "yrTerm1.allUpfront": "3.4374",
            "yrTerm1.partialUpfront": "3.5073",
            "yrTerm3.partialUpfront": "2.4979"
          },
          "ondemand": "6.96"
        }
      },
      "us-west-2": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "3.3",
            "yrTerm3.allUpfront": "1.6462",
            "yrTerm1.allUpfront": "2.7639",
            "yrTerm1.partialUpfront": "2.8202",
            "yrTerm3.partialUpfront": "1.7516"
          },
          "ondemand": "6.198"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "3.216",
            "yrTerm3.allUpfront": "1.5814",
            "yrTerm1.allUpfront": "2.6959",
            "yrTerm1.partialUpfront": "2.7513",
            "yrTerm3.partialUpfront": "1.6823"
          },
          "ondemand": "5.52"
        }
      },
      "us-gov-west-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "5.213",
            "yrTerm3.allUpfront": "3.248",
            "yrTerm1.allUpfront": "4.586",
            "yrTerm1.partialUpfront": "4.652",
            "yrTerm3.partialUpfront": "3.37"
          },
          "ondemand": "7.956"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "3.863",
            "yrTerm3.allUpfront": "1.898",
            "yrTerm1.allUpfront": "3.236",
            "yrTerm1.partialUpfront": "3.302",
            "yrTerm3.partialUpfront": "2.02"
          },
          "ondemand": "6.624"
        }
      },
      "eu-central-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "3.939",
            "yrTerm3.allUpfront": "2.2728",
            "yrTerm1.allUpfront": "3.299",
            "yrTerm1.partialUpfront": "3.3661",
            "yrTerm3.partialUpfront": "2.4179"
          },
          "ondemand": "7.03"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "3.856",
            "yrTerm3.allUpfront": "2.2079",
            "yrTerm1.allUpfront": "3.232",
            "yrTerm1.partialUpfront": "3.2973",
            "yrTerm3.partialUpfront": "2.3489"
          },
          "ondemand": "6.352"
        }
      },
      "eu-west-1": {
        "mswin": {
          "reserved": {
            "yrTerm1.noUpfront": "4.159",
            "yrTerm3.allUpfront": "2.1083",
            "yrTerm1.allUpfront": "3.4838",
            "yrTerm1.partialUpfront": "3.5545",
            "yrTerm3.partialUpfront": "2.2424"
          },
          "ondemand": "6.558"
        },
        "linux": {
          "reserved": {
            "yrTerm1.noUpfront": "3.576",
            "yrTerm3.allUpfront": "2.0435",
            "yrTerm1.allUpfront": "2.9936",
            "yrTerm1.partialUpfront": "3.0549",
            "yrTerm3.partialUpfront": "2.1738"
          },
          "ondemand": "5.88"
        }
      }
    },
    "vpc": {
      "ips_per_eni": 30,
      "max_enis": 8
    },
    "arch": [
      "x86_64"
    ],
    "vpc_only": false,
    "linux_virtualization_types": [
      "HVM"
    ],
    "ebs_optimized": true,
    "storage": {
      "ssd": false,
      "devices": 24,
      "size": 2000.0
    },
    "instance_type": "d2.8xlarge",
    "ECU": 116.0,
    "memory": 244.0,
    "ebs_max_bandwidth": 4000.0
  }
]`
