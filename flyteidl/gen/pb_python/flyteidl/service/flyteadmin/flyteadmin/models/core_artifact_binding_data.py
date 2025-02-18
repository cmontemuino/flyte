# coding: utf-8

"""
    flyteidl/service/admin.proto

    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)  # noqa: E501

    OpenAPI spec version: version not set
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six


class CoreArtifactBindingData(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'index': 'int',
        'partition_key': 'str',
        'transform': 'str'
    }

    attribute_map = {
        'index': 'index',
        'partition_key': 'partition_key',
        'transform': 'transform'
    }

    def __init__(self, index=None, partition_key=None, transform=None):  # noqa: E501
        """CoreArtifactBindingData - a model defined in Swagger"""  # noqa: E501

        self._index = None
        self._partition_key = None
        self._transform = None
        self.discriminator = None

        if index is not None:
            self.index = index
        if partition_key is not None:
            self.partition_key = partition_key
        if transform is not None:
            self.transform = transform

    @property
    def index(self):
        """Gets the index of this CoreArtifactBindingData.  # noqa: E501


        :return: The index of this CoreArtifactBindingData.  # noqa: E501
        :rtype: int
        """
        return self._index

    @index.setter
    def index(self, index):
        """Sets the index of this CoreArtifactBindingData.


        :param index: The index of this CoreArtifactBindingData.  # noqa: E501
        :type: int
        """

        self._index = index

    @property
    def partition_key(self):
        """Gets the partition_key of this CoreArtifactBindingData.  # noqa: E501


        :return: The partition_key of this CoreArtifactBindingData.  # noqa: E501
        :rtype: str
        """
        return self._partition_key

    @partition_key.setter
    def partition_key(self, partition_key):
        """Sets the partition_key of this CoreArtifactBindingData.


        :param partition_key: The partition_key of this CoreArtifactBindingData.  # noqa: E501
        :type: str
        """

        self._partition_key = partition_key

    @property
    def transform(self):
        """Gets the transform of this CoreArtifactBindingData.  # noqa: E501


        :return: The transform of this CoreArtifactBindingData.  # noqa: E501
        :rtype: str
        """
        return self._transform

    @transform.setter
    def transform(self, transform):
        """Sets the transform of this CoreArtifactBindingData.


        :param transform: The transform of this CoreArtifactBindingData.  # noqa: E501
        :type: str
        """

        self._transform = transform

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(CoreArtifactBindingData, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, CoreArtifactBindingData):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
